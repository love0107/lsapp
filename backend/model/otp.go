package model

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type OTP struct {
	Id        int       `orm:"column(id);auto"`
	UserID    int64     `orm:"column(userid)"`
	Mobile    string    `orm:"column(mobile);size(100)"`
	Email     string    `orm:"column(email);size(100)"`
	SentTime  time.Time `orm:"column(senttime);type(datetime);null"`
	ExpireAt  time.Time `orm:"column(expireat);type(datetime);null"`
	Otp       string    `orm:"column(otp);null"`
	Token     string    `orm:"column(token);size(32);null"`
	SentTo    string    `orm:"column(sentto);size(64);null"`
	CreatedOn time.Time `orm:"column(createdon);type(timestamp);auto_now_add"`
	UpdatedOn time.Time `orm:"column(updatedon);type(timestamp);auto_now"`
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

// get the otp by user number
func (otp *OTP) GetOtpByEmail(email string) (*OTP, error) {
	o := orm.NewOrm()
	storedOtp := &OTP{}
	err := o.QueryTable(new(OTP)).Filter("email", email).OrderBy("-id").One(storedOtp)
	if err != nil {
		return nil, err
	}
	return storedOtp, nil
}

func (otp *OTP) GetOtpByUserId(userId int64) (*OTP, error) {
	o := orm.NewOrm()
	storedOtp := &OTP{}
	err := o.QueryTable(new(OTP)).Filter("userId", userId).OrderBy("-id").One(storedOtp)
	if err != nil {
		return nil, err
	}
	return storedOtp, nil
}
