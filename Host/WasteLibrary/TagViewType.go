package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//TagViewType
type TagViewType struct {
	TagId          float64
	ContainerNo    string
	ContainerStatu string
	TagStatu       string
	Latitude       float64
	Longitude      float64
}

//New
func (res *TagViewType) New() {
	res.TagId = 0
	res.ContainerNo = ""
	res.ContainerStatu = CONTAINER_FULLNESS_STATU_NONE
	res.TagStatu = TAG_STATU_NONE
	res.Latitude = 0
	res.Longitude = 0
}

//ToId String
func (res *TagViewType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToTagId String
func (res *TagViewType) ToTagIdString() string {
	return fmt.Sprintf("%.0f", res.TagId)
}

//ToByte
func (res *TagViewType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData

}

//ToString Get JSON
func (res *TagViewType) ToString() string {
	return string(res.ToByte())

}

//Byte To TagViewType
func ByteToTagViewType(retByte []byte) TagViewType {
	var retVal TagViewType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To TagViewType
func StringToTagViewType(retStr string) TagViewType {
	return ByteToTagViewType([]byte(retStr))
}