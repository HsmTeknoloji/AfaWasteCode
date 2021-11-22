package WasteLibrary

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
)

//RecyDeviceAlarmType
type RecyDeviceAlarmType struct {
	DeviceId    float64
	AlarmStatus string
	AlarmTime   string
	AlarmType   string
	Alarm       string
	NewData     bool
}

//New
func (res *RecyDeviceAlarmType) New() {
	res.DeviceId = 0
	res.AlarmStatus = ALARMSTATU_NONE
	res.AlarmTime = GetTime()
	res.AlarmType = ALARMTYPE_NONE
	res.Alarm = ""
	res.NewData = false
}

//GetByRedis
func (res *RecyDeviceAlarmType) GetByRedis() ResultType {
	var resultVal ResultType
	resultVal = GetRedisForStoreApi(REDIS_RECY_ALARM_DEVICES, res.ToIdString())
	if resultVal.Result == RESULT_OK {
		res.StringToType(resultVal.Retval.(string))
		res.NewData = false
	} else {
		return resultVal
	}

	resultVal.Retval = res.ToString()
	return resultVal
}

//SaveToRedis
func (res *RecyDeviceAlarmType) SaveToRedis() ResultType {
	var resultVal ResultType
	resultVal = SaveRedisForStoreApi(REDIS_RECY_ALARM_DEVICES, res.ToIdString(), res.ToString())
	return resultVal
}

//SaveToDb
func (res *RecyDeviceAlarmType) SaveToDb() ResultType {
	var resultVal ResultType
	var currentHttpHeader HttpClientHeaderType
	currentHttpHeader.New()
	currentHttpHeader.DataType = DATATYPE_RECY_ALARM_DEVICE

	data := url.Values{
		HTTP_HEADER: {currentHttpHeader.ToString()},
		HTTP_DATA:   {res.ToString()},
	}
	resultVal = SaveStaticDbMainForStoreApi(data)
	if resultVal.Result == RESULT_OK {
		res.DeviceId = StringIdToFloat64(resultVal.Retval.(string))
		resultVal.Retval = res.ToString()
	}

	return resultVal
}

//ToId String
func (res *RecyDeviceAlarmType) ToIdString() string {
	return fmt.Sprintf("%.0f", res.DeviceId)
}

//ToByte
func (res *RecyDeviceAlarmType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString Get JSON
func (res *RecyDeviceAlarmType) ToString() string {
	return string(res.ToByte())

}

//Byte To RecyDeviceAlarmType
func ByteToRecyDeviceAlarmType(retByte []byte) RecyDeviceAlarmType {
	var retVal RecyDeviceAlarmType
	json.Unmarshal(retByte, &retVal)
	return retVal
}

//String To RecyDeviceAlarmType
func StringToRecyDeviceAlarmType(retStr string) RecyDeviceAlarmType {
	return ByteToRecyDeviceAlarmType([]byte(retStr))
}

//ByteToType
func (res *RecyDeviceAlarmType) ByteToType(retByte []byte) {
	json.Unmarshal(retByte, res)
}

//StringToType
func (res *RecyDeviceAlarmType) StringToType(retStr string) {
	res.ByteToType([]byte(retStr))
}

//SelectSQL
func (res *RecyDeviceAlarmType) SelectSQL() string {
	return fmt.Sprintf(`SELECT AlarmStatus,AlarmTime,AlarmType,Alarm
	 FROM public.recy_alarm_devices
	 WHERE DeviceId=%f ;`, res.DeviceId)
}

//InsertSQL
func (res *RecyDeviceAlarmType) InsertSQL() string {
	return fmt.Sprintf(`INSERT INTO public.recy_alarm_devices (DeviceId,AlarmStatus,AlarmTime,AlarmType,Alarm) 
	  VALUES (%f,'%s','%s','%s','%s') 
	  RETURNING DeviceId;`, res.DeviceId, res.AlarmStatus, res.AlarmTime, res.AlarmType, res.Alarm)
}

//UpdateSQL
func (res *RecyDeviceAlarmType) UpdateSQL() string {
	return fmt.Sprintf(`UPDATE public.recy_alarm_devices 
	  SET AlarmStatus='%s',AlarmTime='%s',AlarmType='%s',Alarm='%s' 
	  WHERE DeviceId=%f  
	  RETURNING DeviceId;`,
		res.AlarmStatus,
		res.AlarmTime,
		res.AlarmType,
		res.Alarm,
		res.DeviceId)
}

//SelectWithDb
func (res *RecyDeviceAlarmType) SelectWithDb(db *sql.DB) error {
	errDb := db.QueryRow(res.SelectSQL()).Scan(
		&res.AlarmStatus,
		&res.AlarmTime,
		&res.AlarmType,
		&res.Alarm)
	return errDb
}
