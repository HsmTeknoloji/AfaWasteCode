package WasteLibrary

import (
	"encoding/json"
)

//CustomersType
type CustomersType struct {
	Customers map[string]float64
}

//New
func (res *CustomersType) New() {
	res.Customers = make(map[string]float64)
}

//GetByRedis
func (res *CustomersType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_CUSTOMERS, REDIS_CUSTOMERS)
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//ToByte
func (res *CustomersType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *CustomersType) ToString() string {
	return string(res.ToByte())

}

//Byte To CustomersType
func ByteToCustomersType(retByte []byte) CustomersType {
	var retVal CustomersType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To CustomersType
func StringToCustomersType(retStr string) CustomersType {
	return ByteToCustomersType([]byte(retStr))
}

//ByteToType
func (res *CustomersType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *CustomersType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}
