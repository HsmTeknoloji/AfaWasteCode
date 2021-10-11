package WasteLibrary

import (
	"encoding/json"
	"fmt"
)

//UserType
type UserType struct {
	UserId     float64
	CustomerId float64
	UserName   string
	UserRole   string
	Password   string
	Email      string
	Active     string
	CreateTime string
}

//ToId String
func (res UserType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.UserId)
}

//ToCustomerId String
func (res UserType) ToCustomerIdString() string {
	return fmt.Sprintf("%.0f", res.CustomerId)
}

//ToByte
func (res UserType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res UserType) ToString() string {
	return string(res.ToByte())

}

//Byte To UserType
func ByteToUserType(retByte []byte) UserType {
	var retVal UserType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To UserType
func StringToUserType(retStr string) UserType {
	return ByteToUserType([]byte(retStr))
}

//SelectSQL
func (res UserType) SelectSQL() string {
	return fmt.Sprintf(`SELECT 
			CustomerId,
			UserName,
			UserRole,
			Password,
			Email,
			Active,
			CreateTime 
			 FROM public.users 
			 WHERE UserId=%f ;`, res.UserId)
}

//InsertSQL
func (res UserType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.users 
	(UserRole,Email,UserName,CustomerId,Password) 
	  VALUES ('%s','%s','%s',%f,'%s')   
	  RETURNING UserId;`,
		res.UserRole, res.Email, res.UserName,
		res.CustomerId, res.Password)
}

//UpdateSQL
func (res UserType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.users 
	SET UserRole='%s',Email='%s',UserName='%s',CustomerId=%f,Active='%s'
	  WHERE UserId=%f  
	RETURNING UserId;`,
		res.UserRole, res.Email, res.UserName,
		res.CustomerId, res.Active, res.UserId)
}

//UpdatePasswordSQL
func (res UserType) UpdatePasswordSQL() string {
	return fmt.Sprintf(`UPDATE public.users 
	SET Password='%s'
	  WHERE UserId=%f  
	RETURNING UserId;`,
		res.Password, res.UserId)
}
