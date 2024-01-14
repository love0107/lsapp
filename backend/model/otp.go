package model

import (
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/text/number"
)

type OTP struct {
	Id       int       `orm:"column(id);auto"`
	Mobile   string    `orm:"column(mobile);size(100)"`
	SentTime time.Time `orm:"column(SentTime);type(datetime);null"`
	Otp      string       `orm:"column(otp);null"`
	Token    string    `orm:"column(token);size(32);null"`
	SentTo   string    `orm:"column(sentTo);size(64);null"`
}

func (otp *OTP) TableName() string {
	return "ls_otp"
}
// add otp and store it
// 
func (otp *OTP)AddOtp()(id int64, err error){
	o:=orm.NewOrm()
	id, err =o.Insert(&otp)
	return id, err
}
// get the otp by user number
func (otp *OTP) GetOtpByNumber(number string)(ls_otp string, err error){
	o:=orm.NewOrm()
	err =o.QueryTable(new(OTP)).Filter("mobile", number).Filter("created_at__gte", startOfDay).OrderBy("-created_at").Limit(1)
	return ls_otp, nil
}