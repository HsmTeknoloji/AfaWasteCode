package main

import (
	"net/http"
	"net/url"

	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
}
func main() {

	initStart()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/reader", reader)
	http.ListenAndServe(":80", nil)
}

func reader(w http.ResponseWriter, req *http.Request) {

	if WasteLibrary.AllowCors {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	}

	var resultVal WasteLibrary.ResultType
	resultVal.Result = WasteLibrary.RESULT_FAIL

	if err := req.ParseForm(); err != nil {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_HTTP_PARSE
		w.Write(resultVal.ToByte())

		WasteLibrary.LogErr(err)
		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.Repeat == WasteLibrary.STATU_PASSIVE {
		if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
			var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
			currentData.DeviceId = currentHttpHeader.DeviceId
			currentData.DeviceGps.DeviceId = currentData.DeviceId
			WasteLibrary.LogStr("Header : " + currentHttpHeader.ToString())
			WasteLibrary.LogStr("Data : " + currentData.ToString())
			currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_GPS_DEVICE
			currentData.DeviceGps.GpsTime = currentHttpHeader.Time
			if currentData.DeviceGps.Longitude != 0 && currentData.DeviceGps.Latitude != 0 {
				data := url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.DeviceGps.ToString()},
				}
				resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}

				currentData.DeviceGps.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))
				data = url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.DeviceGps.ToString()},
				}
				resultVal = WasteLibrary.GetStaticDbMainForStoreApi(data)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_GET
					w.Write(resultVal.ToByte())

					return
				}
				currentData.DeviceGps = WasteLibrary.StringToRfidDeviceGpsType(resultVal.Retval.(string))
				resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_GPS_DEVICES, currentData.DeviceGps.ToIdString(), currentData.DeviceGps.ToString())
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
					w.Write(resultVal.ToByte())

					return
				}
				data = url.Values{
					WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
					WasteLibrary.HTTP_DATA:   {currentData.DeviceGps.ToString()},
				}
				resultVal = WasteLibrary.SaveReaderDbMainForStoreApi(data)
				if resultVal.Result != WasteLibrary.RESULT_OK {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
					w.Write(resultVal.ToByte())

					return
				}
				if currentData.DeviceGps.Speed == 0 {
					//TO DO
					//stop proc
				}

				//TO DO
				//send data web socket
			} else {
				resultVal.Result = WasteLibrary.RESULT_OK
			}
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_OK
	}
	w.Write(resultVal.ToByte())

}
