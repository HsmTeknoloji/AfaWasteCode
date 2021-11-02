package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//RecyDeviceBaseType
type RecyDeviceBaseType struct {
	DeviceId    float64
	ContainerNo string
	DeviceType  string
	NewData     bool
}

//New
func (res *RecyDeviceBaseType) New() {
	res.DeviceId = 0
	res.ContainerNo = ""
	res.DeviceType = RECY_DEVICE_TYPE_NONE
	res.NewData = false
}

//ToId String
func (res RecyDeviceBaseType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res RecyDeviceBaseType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res RecyDeviceBaseType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceBaseType
func ByteToRecyDeviceBaseType(retByte []byte) RecyDeviceBaseType {
	var retVal RecyDeviceBaseType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceBaseType
func StringToRecyDeviceBaseType(retStr string) RecyDeviceBaseType {
	return ByteToRecyDeviceBaseType([]byte(retStr))
}

//SelectSQL
func (res RecyDeviceBaseType) SelectSQL() string {
	return fmt.Sprintf(`SELECT ContainerNo,DeviceType
	 FROM public.recy_base_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res RecyDeviceBaseType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.recy_base_devices (DeviceId,ContainerNo,DeviceType) 
	  VALUES (%f,'%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.ContainerNo, res.DeviceType)
}

//UpdateSQL
func (res RecyDeviceBaseType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_base_devices 
	  SET ContainerNo='%s',DeviceType='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.ContainerNo,
		res.DeviceType,
		res.DeviceId)
}

//SelectWithDb
func (res RecyDeviceBaseType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.ContainerNo,
		&res.DeviceType)
	return errDb
}