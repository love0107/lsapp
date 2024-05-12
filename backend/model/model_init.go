package model

import "github.com/astaxie/beego/orm"

// init all the tables
func InitModel() {
	orm.RegisterModel(new(User))          // user table
	orm.RegisterModel(new(LsConfig))      // config table
	orm.RegisterModel(new(OTP))           // init OTP
	orm.RegisterModel(new(EmailTemplate)) // inti email templates
	orm.RegisterModel(new(SMSTemplate))   // Register model with Beego ORM
}
