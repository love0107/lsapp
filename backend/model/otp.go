package model

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type OTP struct {
	Id       int       `orm:"column(id);auto"`
	UserID   int64     `orm:"column(user_id)"`
	Mobile   string    `orm:"column(mobile);size(100)"`
	SentTime time.Time `orm:"column(sentTime);type(datetime);null"`
	ExpireAt time.Time `orm:"column(expireAt);type(datetime);null"`
	Otp      string    `orm:"column(otp);null"`
	Token    string    `orm:"column(token);size(32);null"`
	SentTo   string    `orm:"column(sentTo);size(64);null"`
}

func (otp *OTP) TableName() string {
	return "ls_otp"
}

// add otp and store it
func (otp *OTP) AddOtp() (id int64, err error) {
	o := orm.NewOrm()
	fmt.Println("otp----->", &otp)
	id, err = o.Insert(otp)
	if err != nil {
		fmt.Println("Error inserting OTP:", err)
	}
	return id, err
}

// // get the otp by user number
// func (otp *OTP) GetOtpByNumber(number string) (ls_otp string, err error) {
// 	o := orm.NewOrm()
// 	err = o.QueryTable(new(OTP)).Filter("mobile", number).Filter("created_at__gte", startOfDay).OrderBy("-created_at").Limit(1).One(&ls_otp)
// 	return ls_otp, nil
// }
