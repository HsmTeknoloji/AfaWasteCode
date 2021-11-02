package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//UltDeviceGpsType
type UltDeviceGpsType struct {
	DeviceId  float64
	Latitude  float64
	Longitude float64
	GpsTime   string
	NewData   bool
}

//New
func (res *UltDeviceGpsType) New() {
	res.DeviceId = 0
	res.Latitude = 0
	res.Longitude = 0
	res.GpsTime = ""
	res.NewData = false
}

//ToId String
func (res UltDeviceGpsType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res UltDeviceGpsType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res UltDeviceGpsType) ToString() string {
	return string(res.ToByte())

}

//Byte To UltDeviceGpsType
func ByteToUltDeviceGpsType(retByte []byte) UltDeviceGpsType {
	var retVal UltDeviceGpsType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UltDeviceGpsType
func StringToUltDeviceGpsType(retStr string) UltDeviceGpsType {
	return ByteToUltDeviceGpsType([]byte(retStr))
}

//SelectSQL
func (res UltDeviceGpsType) SelectSQL() string {
	return fmt.Sprintf(`SELECT Latitude,Longitude,GpsTime
	 FROM public.ult_gps_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res UltDeviceGpsType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.ult_gps_devices (DeviceId,Latitude,Longitude,GpsTime) 
	  VALUES (%f,%f,%f,'%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.Latitude, res.Longitude, res.GpsTime)
}

//UpdateSQL
func (res UltDeviceGpsType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.ult_gps_devices 
	  SET Latitude=%f,Longitude=%f,GpsTime='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.Latitude,
		res.Longitude,
		res.GpsTime,
		res.DeviceId)
}

//SelectWithDb
func (res UltDeviceGpsType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.Latitude,
		&res.Longitude,
		&res.GpsTime)
	return errDb
}