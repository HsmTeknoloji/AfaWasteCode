package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//UltDeviceViewType
type UltDeviceViewType struct {
	DeviceId       float64
	ContainerNo    string
	ContainerStatu string
	UltStatus      string
	Latitude       float64
	Longitude      float64
	SensPercent    float64
}

//New
func (res *UltDeviceViewType) New() {
	res.DeviceId = 0
	res.ContainerNo = ""
	res.ContainerStatu = CONTAINER_FULLNESS_STATU_NONE
	res.UltStatus = ULT_STATU_NONE
	res.Latitude = 0
	res.Longitude = 0
	res.SensPercent = 0
}

//ToId String
func (res *UltDeviceViewType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToDeviceId String
func (res *UltDeviceViewType) ToDeviceIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *UltDeviceViewType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *UltDeviceViewType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceViewType
func ByteToUltDeviceViewType(retByte []byte) UltDeviceViewType {
	var retVal UltDeviceViewType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceViewType
func StringToUltDeviceViewType(retStr string) UltDeviceViewType {
	return ByteToUltDeviceViewType([]byte(retStr))
}