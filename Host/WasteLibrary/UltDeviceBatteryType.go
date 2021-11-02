package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//UltDeviceBatteryType
type UltDeviceBatteryType struct {
	DeviceId      float64
	Battery       string
	BatteryStatus string
	BatteryTime   string
	NewData       bool
}

//New
func (res *UltDeviceBatteryType) New() {
	res.DeviceId = 0
	res.Battery = "0000"
	res.BatteryStatus = BATTERYSTATU_NONE
	res.BatteryTime = ""
	res.NewData = false
}

//ToId String
func (res UltDeviceBatteryType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res UltDeviceBatteryType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res UltDeviceBatteryType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceBatteryType
func ByteToUltDeviceBatteryType(retByte []byte) UltDeviceBatteryType {
	var retVal UltDeviceBatteryType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceBatteryType
func StringToUltDeviceBatteryType(retStr string) UltDeviceBatteryType {
	return ByteToUltDeviceBatteryType([]byte(retStr))
}

//SelectSQL
func (res UltDeviceBatteryType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Battery,BatteryStatus,BatteryTime
	 FROM public.ult_battery_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res UltDeviceBatteryType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_battery_devices (DeviceId,Battery,BatteryStatus,BatteryTime) 
	  VALUES (%f,'%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Battery, res.BatteryStatus, res.BatteryTime)
}

//UpdateSQL
func (res UltDeviceBatteryType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_battery_devices 
	  SET Battery='%s',BatteryStatus='%s',BatteryTime='%s'
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Battery,
		res.BatteryStatus,
		res.BatteryTime,
		res.DeviceId)
}

//SelectWithDb
func (res UltDeviceBatteryType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Battery,
		&res.BatteryStatus,
		&res.BatteryTime)
	return errDb
}