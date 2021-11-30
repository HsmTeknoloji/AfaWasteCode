package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//RfidDeviceType
type RfidDeviceType struct {
	DeviceId         float64
	DeviceMain       RfidDeviceMainType
	DeviceBase       RfidDeviceBaseType
	DeviceStatu      RfidDeviceStatuType
	DeviceGps        RfidDeviceGpsType
	DeviceAlarm      RfidDeviceAlarmType
	DeviceTherm      RfidDeviceThermType
	DeviceVersion    RfidDeviceVersionType
	DeviceDetail     RfidDeviceDetailType
	DeviceWorkHour   RfidDeviceWorkHourType
	DeviceEmbededGps RfidDeviceEmbededGpsType
}

//New
func (res *RfidDeviceType) New() {
	res.DeviceId = 0
	res.DeviceMain.New()
	res.DeviceBase.New()
	res.DeviceGps.New()
	res.DeviceTherm.New()
	res.DeviceVersion.New()
	res.DeviceAlarm.New()
	res.DeviceStatu.New()
	res.DeviceDetail.New()
	res.DeviceWorkHour.New()
	res.DeviceEmbededGps.New()
}

//GetByRedis
func (res *RfidDeviceType) GetByRedis(dbIndex string) ResultType {
	var resultVal ResultType

	res.DeviceMain.DeviceId = res.DeviceId
	resultVal = res.DeviceMain.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceBase.DeviceId = res.DeviceId
	resultVal = res.DeviceBase.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceGps.DeviceId = res.DeviceId
	resultVal = res.DeviceGps.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceEmbededGps.DeviceId = res.DeviceId
	resultVal = res.DeviceEmbededGps.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceTherm.DeviceId = res.DeviceId
	resultVal = res.DeviceTherm.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceVersion.DeviceId = res.DeviceId
	resultVal = res.DeviceVersion.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceAlarm.DeviceId = res.DeviceId
	resultVal = res.DeviceAlarm.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceStatu.DeviceId = res.DeviceId
	resultVal = res.DeviceStatu.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceDetail.DeviceId = res.DeviceId
	resultVal = res.DeviceDetail.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}
	res.DeviceWorkHour.DeviceId = res.DeviceId
	resultVal = res.DeviceWorkHour.GetByRedis(dbIndex)
	if resultVal.Result != RESULT_OK {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//GetByRedisBySerial
func (res *RfidDeviceType) GetByRedisBySerial(serial string) ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi("0", REDIS_SERIAL_RFID_DEVICE, serial)
	if resultVal.Result == RESULT_OK {
		var deviceId string = resultVal.Retval.(string)
		res.DeviceId = StringIdToFloat64(deviceId)
		resultVal = res.GetByRedis("0")
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//ToId String
func (res *RfidDeviceType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceType
func ByteToRfidDeviceType(retByte []byte) RfidDeviceType {
	var retVal RfidDeviceType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceType
func StringToRfidDeviceType(retStr string) RfidDeviceType {
	return ByteToRfidDeviceType([]byte(retStr))
}
