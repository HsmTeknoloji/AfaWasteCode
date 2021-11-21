package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//AdminConfigType
type AdminConfigType struct {
	CustomerId      float64
	DeviceBaseWork  string
	WorkStartHour   int
	WorkStartMinute int
	WorkEndHour     int
	WorkEndMinute   int
	Active          string
	CreateTime      string
}

//New
func (res *AdminConfigType) New() {
	res.CustomerId = 1
	res.WorkStartHour = 06
	res.WorkStartMinute = 0
	res.WorkEndHour = 18
	res.WorkEndMinute = 30
	res.DeviceBaseWork = STATU_PASSIVE
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//GetByRedis
func (res *AdminConfigType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_CUSTOMER_ADMINCONFIG, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//ToId String
func (res *AdminConfigType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res *AdminConfigType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *AdminConfigType) ToString() string {
	return string(res.ToByte())

}

//Byte To AdminConfigType
func ByteToAdminConfigType(retByte []byte) AdminConfigType {
	var retVal AdminConfigType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To AdminConfigType
func StringToAdminConfigType(retStr string) AdminConfigType {
	return ByteToAdminConfigType([]byte(retStr))
}

//ByteToType
func (res *AdminConfigType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *AdminConfigType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
