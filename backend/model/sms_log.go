package model


type SmsLog struct{
Code int
Status string
Body string
}

func (t *SmsLog) TableName() string {
	return "ls_sms_log"
}

// func (t *SmsLog) AddSmsLog() (int64, error) {
// 	return DB.Insert(t)
// }
