package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

//TagMainType
type TagMainType struct {
	TagId      float64
	CustomerId float64
	DeviceId   float64
	Epc        string
	Active     string
	CreateTime string
}

//New
func (res *TagMainType) New() {
	res.TagId = 0
	res.CustomerId = 1
	res.TagId = 0
	res.Epc = ""
	res.Active = STATU_ACTIVE
	res.CreateTime = GetTime()
}

//ToId String
func (res *TagMainType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToCustomerId String
func (res *TagMainType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToTagId String
func (res *TagMainType) ToTagIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagMainType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagMainType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagMainType
func ByteToTagMainType(retByte []byte) TagMainType {
	var retVal TagMainType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagMainType
func StringToTagMainType(retStr string) TagMainType {
	return ByteToTagMainType([]byte(retStr))
}

//SelectSQL
func (res *TagMainType) SelectSQL() string {
	return fmt.Sprintf(`SELECT CustomerId,DeviceId,Epc,Active,CreateTime
	 FROM public.tag_mains
	 WHERE TagId=%f  ;`, res.TagId)
}

//InsertSQL
func (res *TagMainType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tag_mains (CustomerId,DeviceId,Epc) 
	  VALUES (%f,%f,'%s') 
	  RETURNING TagId;`, res.CustomerId, res.DeviceId, res.Epc)
}

//InsertDataSQL
func (res *TagMainType) InsertDataSQL() string {
	return fmt.Sprintf(`INSERT INTO public.tag_mains (TagId,CustomerId,DeviceId,Epc) 
	  VALUES (%f,%f,%f,'%s') 
	  RETURNING TagId;`, res.TagId, res.CustomerId, res.DeviceId, res.Epc)
}

//UpdateSQL
func (res *TagMainType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.tag_mains 
	  SET CustomerId=%f,DeviceId=%f,Epc='%s' 
	  WHERE TagId=%f  
	  RETURNING TagId;`,
		res.CustomerId,
		res.DeviceId,
		res.Epc,
		res.TagId)
}

//SelectWithDb
func (res *TagMainType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.CustomerId,
		&res.DeviceId,
		&res.Epc,
		&res.Active,
		&res.CreateTime)
	return errDb
}