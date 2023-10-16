package model

import "github.com/beego/beego/orm"

func InitModel() {
	orm.RegisterModel(new(User))// user table
    orm.RegisterModel(new(LsConfig))// config table 
}
