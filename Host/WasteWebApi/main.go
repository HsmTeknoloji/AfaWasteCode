package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/devafatek/WasteLibrary"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	go WasteLibrary.InitLog()
}
func main() {

	initStart()

	http.HandleFunc("/health", WasteLibrary.HealthHandler)
	http.HandleFunc("/readiness", WasteLibrary.ReadinessHandler)
	http.HandleFunc("/status", WasteLibrary.StatusHandler)
	http.HandleFunc("/getCustomer", getCustomer)
	http.HandleFunc("/getDevice", getDevice)
	http.HandleFunc("/getDevices", getDevices)
	http.HandleFunc("/getConfig", getConfig)
	http.HandleFunc("/getTags", getTags)
	http.HandleFunc("/getTag", getTag)
	http.HandleFunc("/setDevice", setDevice)
	http.HandleFunc("/setUser", setUser)
	http.HandleFunc("/setConfig", setConfig)
	http.HandleFunc("/getUser", getUser)
	http.HandleFunc("/getUsers", getUsers)

	http.ListenAndServe(":80", nil)
}

func getCustomer(w http.ResponseWriter, req *http.Request) {

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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}
	w.Write(resultVal.ToByte())

}

func getDevice(w http.ResponseWriter, req *http.Request) {

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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
		if currentData.DeviceMain.CustomerId == linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentData.ToString()
			w.Write(resultVal.ToByte())

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentData WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
		if currentData.DeviceMain.CustomerId == linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentData.ToString()
			w.Write(resultVal.ToByte())

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

		}
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var currentData WasteLibrary.RecyDeviceType = WasteLibrary.StringToRecyDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		currentData.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
		if currentData.DeviceMain.CustomerId == linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_OK
			resultVal.Retval = currentData.ToString()
			w.Write(resultVal.ToByte())

		} else {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
		w.Write(resultVal.ToByte())

	}

}

func getDevices(w http.ResponseWriter, req *http.Request) {

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

	fmt.Println(req.Form)

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}
	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))

	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var customerDevices WasteLibrary.CustomerRfidDevicesType
		customerDevices.CustomerId = linkCustomer.CustomerId
		resultVal = customerDevices.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevicesList WasteLibrary.CustomerRfidDevicesListType
		customerDevicesList.New()
		customerDevicesList.CustomerId = linkCustomer.CustomerId
		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {

				var currentDevice WasteLibrary.RfidDeviceType
				currentDevice.New()
				currentDevice.DeviceId = deviceId
				resultVal = currentDevice.GetByRedis()
				if resultVal.Result == WasteLibrary.RESULT_OK {
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}

			}
		}
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var customerDevices WasteLibrary.CustomerUltDevicesType
		customerDevices.CustomerId = linkCustomer.CustomerId
		resultVal = customerDevices.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevicesList WasteLibrary.CustomerUltDevicesViewListType
		customerDevicesList.New()
		customerDevicesList.CustomerId = linkCustomer.CustomerId
		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {
				var currentDevice WasteLibrary.UltDeviceType
				currentDevice.New()
				currentDevice.DeviceId = deviceId
				resultVal = currentDevice.GetByRedis()

				if resultVal.Result == WasteLibrary.RESULT_OK {

					var currentUltDeviceView WasteLibrary.UltDeviceViewType
					currentUltDeviceView.New()
					currentUltDeviceView.DeviceId = currentDevice.DeviceId
					currentUltDeviceView.ContainerNo = currentDevice.DeviceBase.ContainerNo
					currentUltDeviceView.ContainerStatu = currentDevice.DeviceStatu.ContainerStatu
					currentUltDeviceView.UltStatus = currentDevice.DeviceStatu.UltStatus
					currentUltDeviceView.Latitude = currentDevice.DeviceGps.Latitude
					currentUltDeviceView.Longitude = currentDevice.DeviceGps.Longitude
					currentUltDeviceView.SensPercent = currentDevice.DeviceStatu.SensPercent

					customerDevicesList.Devices[currentUltDeviceView.ToIdString()] = currentUltDeviceView
				}

			}
		}
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var customerDevices WasteLibrary.CustomerRecyDevicesType
		customerDevices.CustomerId = linkCustomer.CustomerId
		resultVal = customerDevices.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		var customerDevicesList WasteLibrary.CustomerRecyDevicesListType
		customerDevicesList.New()
		customerDevicesList.CustomerId = linkCustomer.CustomerId

		for _, deviceId := range customerDevices.Devices {

			if deviceId != 0 {
				var currentDevice WasteLibrary.RecyDeviceType
				currentDevice.New()
				currentDevice.DeviceId = deviceId
				resultVal = currentDevice.GetByRedis()
				if resultVal.Result == WasteLibrary.RESULT_OK {
					customerDevicesList.Devices[currentDevice.ToIdString()] = currentDevice
				}

			}
		}
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = customerDevicesList.ToString()
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
	}
	w.Write(resultVal.ToByte())

}

func getConfig(w http.ResponseWriter, req *http.Request) {

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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ADMINCONFIG {
		var customerConfig WasteLibrary.AdminConfigType
		customerConfig.CustomerId = linkCustomer.CustomerId
		resultVal = customerConfig.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_ADMINCONFIG_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_LOCALCONFIG {
		var customerConfig WasteLibrary.LocalConfigType
		customerConfig.CustomerId = linkCustomer.CustomerId
		resultVal = customerConfig.GetByRedis()
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_LOCALCONFIG_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}
	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
	}

	w.Write(resultVal.ToByte())

}

func getTags(w http.ResponseWriter, req *http.Request) {

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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var customerTagsList WasteLibrary.CustomerTagsViewListType
	customerTagsList.CustomerId = linkCustomer.CustomerId
	resultVal = customerTagsList.GetByRedisByReel()

	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = customerTagsList.ToString()
	w.Write(resultVal.ToByte())

}

func getTag(w http.ResponseWriter, req *http.Request) {

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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var currentData WasteLibrary.TagType = WasteLibrary.StringToTagType(req.FormValue(WasteLibrary.HTTP_DATA))

	resultVal = currentData.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_TAG_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}
	currentData = WasteLibrary.StringToTagType(resultVal.Retval.(string))
	currentData.GetByRedis()
	if currentData.TagMain.CustomerId == linkCustomer.CustomerId {
		resultVal.Result = WasteLibrary.RESULT_OK
		resultVal.Retval = currentData.ToString()
		w.Write(resultVal.ToByte())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_TAG_NOTFOUND
		w.Write(resultVal.ToByte())

	}
}

func setDevice(w http.ResponseWriter, req *http.Request) {

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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RFID {
		var currentData WasteLibrary.RfidDeviceType = WasteLibrary.StringToRfidDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.RfidDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetByRedis()
		currentData.DeviceMain.CustomerId = currentOldData.DeviceMain.CustomerId
		if currentData.DeviceMain.CustomerId != linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())
			return
		}

		//DeviceBase
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_BASE_DEVICE
		currentData.DeviceBase.DeviceId = currentData.DeviceId
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.DeviceBase.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentData.DeviceBase.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_BASE_DEVICES, currentData.DeviceBase.ToIdString(), currentData.DeviceBase.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_BASE_DEVICE, currentData.DeviceBase.ToString())

		//DeviceDetail
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RFID_DETAIL_DEVICE
		currentData.DeviceDetail.DeviceId = currentData.DeviceId
		data = url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.DeviceDetail.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentData.DeviceDetail.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RFID_DETAIL_DEVICES, currentData.DeviceDetail.ToIdString(), currentData.DeviceDetail.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RFID_DETAIL_DEVICE, currentData.DeviceDetail.ToString())

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_ULT {
		var currentData WasteLibrary.UltDeviceType = WasteLibrary.StringToUltDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.UltDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetByRedis()
		currentData.DeviceMain.CustomerId = currentOldData.DeviceMain.CustomerId
		if currentData.DeviceMain.CustomerId != linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceBase
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_ULT_BASE_DEVICE
		currentData.DeviceBase.DeviceId = currentData.DeviceId
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.DeviceBase.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentData.DeviceBase.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_ULT_BASE_DEVICES, currentData.DeviceBase.ToIdString(), currentData.DeviceBase.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_ULT_BASE_DEVICE, currentData.DeviceBase.ToString())

	} else if currentHttpHeader.DeviceType == WasteLibrary.DEVICETYPE_RECY {
		var currentData WasteLibrary.RecyDeviceType = WasteLibrary.StringToRecyDeviceType(req.FormValue(WasteLibrary.HTTP_DATA))
		var currentOldData WasteLibrary.RecyDeviceType
		currentOldData.DeviceId = currentData.DeviceId
		currentOldData.GetByRedis()
		currentData.DeviceMain.CustomerId = currentOldData.DeviceMain.CustomerId
		if currentData.DeviceMain.CustomerId != linkCustomer.CustomerId {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
			w.Write(resultVal.ToByte())

			return
		}

		//DeviceBase
		currentHttpHeader.DataType = WasteLibrary.DATATYPE_RECY_BASE_DEVICE
		currentData.DeviceBase.DeviceId = currentData.DeviceId
		data := url.Values{
			WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
			WasteLibrary.HTTP_DATA:   {currentData.DeviceBase.ToString()},
		}
		resultVal = WasteLibrary.SaveStaticDbMainForStoreApi(data)
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		currentData.DeviceBase.DeviceId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_RECY_BASE_DEVICES, currentData.DeviceBase.ToIdString(), currentData.DeviceBase.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.DeviceMain.ToCustomerIdString(), WasteLibrary.DATATYPE_RECY_BASE_DEVICE, currentData.DeviceBase.ToString())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DEVICE_NOTFOUND
	}

	w.Write(resultVal.ToByte())

}

func setConfig(w http.ResponseWriter, req *http.Request) {

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
	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType = WasteLibrary.StringToHttpClientHeaderType(req.FormValue(WasteLibrary.HTTP_HEADER))
	if currentHttpHeader.DataType == WasteLibrary.DATATYPE_ADMINCONFIG {
		var currentData WasteLibrary.AdminConfigType = WasteLibrary.StringToAdminConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_ADMINCONFIG, linkCustomer.ToIdString(), currentData.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())

			return
		}
		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.ToIdString(), WasteLibrary.DATATYPE_ADMINCONFIG, currentData.ToString())
	} else if currentHttpHeader.DataType == WasteLibrary.DATATYPE_LOCALCONFIG {
		var currentData WasteLibrary.LocalConfigType = WasteLibrary.StringToLocalConfigType(req.FormValue(WasteLibrary.HTTP_DATA))
		resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_LOCALCONFIG, linkCustomer.ToIdString(), currentData.ToString())
		if resultVal.Result != WasteLibrary.RESULT_OK {
			resultVal.Result = WasteLibrary.RESULT_FAIL
			resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
			w.Write(resultVal.ToByte())
			return
		}

		resultVal = WasteLibrary.PublishRedisForStoreApi(WasteLibrary.REDIS_CUSTOMER_CHANNEL+currentData.ToIdString(), WasteLibrary.DATATYPE_LOCALCONFIG, currentData.ToString())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DATATYPE
		w.Write(resultVal.ToByte())

		return
	}

	w.Write(resultVal.ToByte())

}

func setUser(w http.ResponseWriter, req *http.Request) {

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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
	var currentDbData WasteLibrary.UserType
	currentDbData.UserId = currentData.UserId
	resultVal = currentDbData.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}
	if currentDbData.CustomerId != linkCustomer.CustomerId {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var currentCustomerUsers WasteLibrary.CustomerUsersType
	currentCustomerUsers.CustomerId = linkCustomer.CustomerId
	resultVal = currentCustomerUsers.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	for _, userId := range currentCustomerUsers.Users {
		if userId != 0 && userId != currentData.UserId {
			var inRedisUser WasteLibrary.UserType
			inRedisUser.UserId = userId
			resultVal = inRedisUser.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK {

				if inRedisUser.UserName == currentData.UserName {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_USERNAMEEXIST
					w.Write(resultVal.ToByte())

					return
				}

				if inRedisUser.Email == currentData.Email && inRedisUser.Email != "" {
					resultVal.Result = WasteLibrary.RESULT_FAIL
					resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_USEREMAILEXIST
					w.Write(resultVal.ToByte())

					return
				}
			}
		}
	}

	var currentHttpHeader WasteLibrary.HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.CustomerId = linkCustomer.CustomerId
	currentHttpHeader.DataType = WasteLibrary.DATATYPE_USER
	data := url.Values{
		WasteLibrary.HTTP_HEADER: {currentHttpHeader.ToString()},
		WasteLibrary.HTTP_DATA:   {currentData.ToString()},
	}
	resultVal = WasteLibrary.SaveConfigDbMainForStoreApi(data)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_DB_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	currentData.UserId = WasteLibrary.StringIdToFloat64(resultVal.Retval.(string))

	resultVal = WasteLibrary.SaveRedisForStoreApi(WasteLibrary.REDIS_USERS, currentData.ToIdString(), currentData.ToString())
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_REDIS_SAVE
		w.Write(resultVal.ToByte())

		return
	}
	w.Write(resultVal.ToByte())

}

func getUser(w http.ResponseWriter, req *http.Request) {

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

	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}
	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var currentData WasteLibrary.UserType = WasteLibrary.StringToUserType(req.FormValue(WasteLibrary.HTTP_DATA))
	resultVal = currentData.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}
	if currentData.CustomerId == linkCustomer.CustomerId {
		w.Write(resultVal.ToByte())

	} else {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_USER_NOTFOUND
		w.Write(resultVal.ToByte())

	}
}

func getUsers(w http.ResponseWriter, req *http.Request) {

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
	var linkCustomer WasteLibrary.CustomerType
	resultVal = linkCustomer.GetByRedisByLink(req.Host)
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	resultVal = checkAuth(req.Form, linkCustomer.ToIdString())

	if resultVal.Result != WasteLibrary.RESULT_OK {
		w.Write(resultVal.ToByte())

		return
	}

	var customerUsers WasteLibrary.CustomerUsersType
	customerUsers.CustomerId = linkCustomer.CustomerId
	resultVal = customerUsers.GetByRedis()
	if resultVal.Result != WasteLibrary.RESULT_OK {
		resultVal.Result = WasteLibrary.RESULT_FAIL
		resultVal.Retval = WasteLibrary.RESULT_ERROR_CUSTOMER_NOTFOUND
		w.Write(resultVal.ToByte())

		return
	}

	var customerUsersList WasteLibrary.CustomerUsersListType = WasteLibrary.CustomerUsersListType{
		CustomerId: linkCustomer.CustomerId,
		Users:      make(map[string]WasteLibrary.UserType),
	}
	for _, userId := range customerUsers.Users {

		if userId != 0 {
			var currentUser WasteLibrary.UserType
			currentUser.UserId = userId
			resultVal = currentUser.GetByRedis()
			if resultVal.Result == WasteLibrary.RESULT_OK {

				customerUsersList.Users[currentUser.ToIdString()] = currentUser
			}

		}
	}
	resultVal.Result = WasteLibrary.RESULT_OK
	resultVal.Retval = customerUsersList.ToString()

	w.Write(resultVal.ToByte())

}

func checkAuth(data url.Values, customerId string) WasteLibrary.ResultType {
	return WasteLibrary.CheckAuth(data, customerId, "ADMIN")

}
