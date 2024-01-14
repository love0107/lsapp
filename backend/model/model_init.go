package model

import (
	"github.com/beego/beego/orm"
)

// init all the tables
func InitModel() {
	orm.RegisterModel(new(User))     // user table
	orm.RegisterModel(new(LsConfig)) // config table
	orm.RegisterModel(new(OTP))      // init OTP
}
