package otp

import "github.com/astaxie/beego/orm"

type OTP struct {
	Id     int    `json:"id"`
	Otp    string `json:"otp"`
	Reason string `json:"reason"`
	UserId int    `json:"userId"`
}

func (otp *OTP) TableName() string {
	return "ls_otp"
}
func RegisterModels() {
	orm.RegisterModel(new(OTP))
}
