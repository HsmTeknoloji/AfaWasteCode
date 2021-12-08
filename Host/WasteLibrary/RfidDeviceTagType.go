package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//RfidDeviceTagType
type RfidDeviceTagType struct {
	DeviceId  float64
	EmptyTags []TagViewType
	FullTags  []TagViewType
}

//New
func (res *RfidDeviceTagType) New() {
	res.DeviceId = 0
	res.EmptyTags = []TagViewType{}
	res.FullTags = []TagViewType{}
}

//GetByRedis
func (res *RfidDeviceTagType) GetByRedis(dbIndex string) ResultType {

	var resultVal ResultType
	resultVal = GetRedisForStoreApi(dbIndex, REDIS_RFID_TAGS, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *RfidDeviceTagType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RFID_TAGS, res.ToIdString(), res.ToString())
	return resultVal
}

//ToId String
func (res *RfidDeviceTagType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RfidDeviceTagType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RfidDeviceTagType) ToString() string {
	return string(res.ToByte())

}

//Byte To RfidDeviceTagType
func ByteToRfidDeviceTagType(retByte []byte) RfidDeviceTagType {
	var retVal RfidDeviceTagType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RfidDeviceTagType
func StringToRfidDeviceTagType(retStr string) RfidDeviceTagType {
	return ByteToRfidDeviceTagType([]byte(retStr))
}

//ByteToType
func (res *RfidDeviceTagType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RfidDeviceTagType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}