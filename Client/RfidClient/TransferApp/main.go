package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/devafatek/WasteLibrary"
)

var applicationType = WasteLibrary.APPTYPE_RFID
var serialNumber = "0"
var currentUser string
var opInterval time.Duration = 5 * 60
var wg sync.WaitGroup

const (
	AWS_S3_REGION = "eu-central-1"
	AWS_S3_BUCKET = "afatek-waste-videos-s3"
)

func initStart() {

	time.Sleep(5 * time.Second)
	WasteLibrary.LogStr("Successfully connected!")
	WasteLibrary.Version = "1"
	WasteLibrary.LogStr("Version : " + WasteLibrary.Version)
	currentUser = WasteLibrary.GetCurrentUser()
	serialNumber = getSerialNumber()
	WasteLibrary.LogStr(currentUser)
	WasteLibrary.LogStr(serialNumber)
}
func main() {

	initStart()

	time.Sleep(time.Second)
	go fileCheck(WasteLibrary.READERTYPE_RF)
	wg.Add(1)

	time.Sleep(time.Second)
	go fileCheck(WasteLibrary.READERTYPE_CAM)
	wg.Add(1)

	time.Sleep(time.Second)
	go fileCheck(WasteLibrary.READERTYPE_GPS)
	wg.Add(1)

	time.Sleep(time.Second)
	go fileCheck(WasteLibrary.READERTYPE_THERM)
	wg.Add(1)

	time.Sleep(time.Second)
	go fileCheck(WasteLibrary.READERTYPE_STATUS)
	wg.Add(1)

	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/trans", trans)
	http.ListenAndServe(":10000", nil)

	wg.Wait()

}

func trans(w http.ResponseWriter, req *http.Request) {
	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	var resultVal WasteLibrary.ResultType

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())
		WasteLibrary.LogErr(err)
		return

	} else {

		readerType := req.FormValue(WasteLibrary.HTTP_READERTYPE)
		dataVal := req.FormValue(WasteLibrary.HTTP_DATA)
		resultVal = sendDataToServer(readerType, dataVal, WasteLibrary.GetTime(), WasteLibrary.STATU_PASSIVE)
		WasteLibrary.LogStr("Send Data To Server : " + resultVal.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			if readerType != WasteLibrary.READERTYPE_CAM {
				storeData(readerType, dataVal)
			}
		}
		if readerType == WasteLibrary.READERTYPE_CAM {
			var curretnTagType WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))

			sendFileToServer(curretnTagType.TagReader.UID)
		}
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())
}

func sendFileToServer(fileName string) {
	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		WasteLibrary.LogErr(err)
	} else {
		err = uploadFile(session, "WAIT_CAM/"+fileName+".mp4")
		if err != nil {
			WasteLibrary.LogErr(err)
		} else {
			WasteLibrary.RemoveFile("WAIT_CAM/" + fileName + ".mp4")
		}

		err = uploadFile(session, "WAIT_CAM/"+fileName+".png")
		if err != nil {
			WasteLibrary.LogErr(err)
		} else {
			WasteLibrary.RemoveFile("WAIT_CAM/" + fileName + ".png")
		}
	}
}

func uploadFile(session *session.Session, uploadFileDir string) error {

	upFile, err := os.Open(uploadFileDir)
	if err != nil {
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(AWS_S3_BUCKET),
		Key:                  aws.String(uploadFileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}

func sendDataToServer(readerType string, sendData string, dataTime string, repeat string) WasteLibrary.ResultType {
	var resultVal WasteLibrary.ResultType
	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.AppType = applicationType
	currentHttpHeader.DeviceNo = serialNumber
	currentHttpHeader.ReaderType = readerType
	currentHttpHeader.Time = dataTime
	currentHttpHeader.Repeat = repeat
	currentHttpHeader.DeviceType = WasteLibrary.DEVICETYPE_RFID
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {sendData},
	}
	resultVal = WasteLibrary.HttpPostReq("http://listener.aws.afatek.com.tr/data", data)
	return resultVal
}

func storeData(dataType string, sendData string) {
	err := ioutil.WriteFile("WAIT_"+dataType+"/"+WasteLibrary.GetTime(), []byte(sendData), 0644)
	if err != nil {
		WasteLibrary.LogErr(err)
	}
}

func resendData(readerType string, fileName string) {
	var resultVal WasteLibrary.ResultType
	if readerType == WasteLibrary.READERTYPE_CAM {
		sendFileToServer(fileName)
	} else {

		var dataJSON string = ""
		var dataTime string = fileName

		readByte, err := ioutil.ReadFile("WAIT_" + readerType + "/" + fileName)
		if err != nil {
			WasteLibrary.LogErr(err)
		} else {

			dataJSON = string(readByte)

			WasteLibrary.LogStr("Read File : " + dataJSON)

			resultVal = sendDataToServer(readerType, string(dataJSON), dataTime, WasteLibrary.STATU_ACTIVE)
			WasteLibrary.LogStr("Send Data To Server Again : " + resultVal.ToString())
			if resultVal.Result == WasteLibrary.RESULT_OK {
				WasteLibrary.RemoveFile("WAIT_" + readerType + "/" + fileName)
			}
		}
	}
}

func fileCheck(readerType string) {
	WasteLibrary.LogStr("File Check :" + readerType)
	for {
		time.Sleep(opInterval * time.Second)

		f, err := os.Open("WAIT_" + readerType)
		if err != nil {
			WasteLibrary.LogErr(err)
			continue
		}
		fileInfo, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			WasteLibrary.LogErr(err)
			continue
		}

		for _, file := range fileInfo {
			time.Sleep(time.Second)
			second := time.Since(file.ModTime()).Seconds()
			if second > 60*60 && second < 24*60*60 {
				var fileName string = file.Name()
				if readerType == WasteLibrary.READERTYPE_CAM {
					spData := strings.Split(strings.TrimSpace(file.Name()), ".")
					fileName = spData[0]
				}
				resendData(readerType, fileName)
			}
			if second > 24*60*60 {
				WasteLibrary.RemoveFile("WAIT_" + readerType + "/" + file.Name())
			}
		}
	}

	wg.Done()
}

func getSerialNumber() string {
	var tempNumber string = ""
	out, err := exec.Command("/home/pi/getSerialNumber.sh").Output()
	if err != nil {
		WasteLibrary.LogErr(err)
	}
	tempNumber = strings.TrimSuffix(string(out), "\n")

	return tempNumber
}