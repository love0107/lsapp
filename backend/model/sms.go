package model

import "time"

// LSSMSTemplate represents the ls_sms_template table in the database
type SMSTemplate struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name)"`
	Content    string    `orm:"column(content)"`
	CreatedOn  time.Time `orm:"column(created_on);type(timestamp);auto_now_add"`
	UpdatedOn  time.Time `orm:"column(updated_on);type(timestamp);auto_now"`
	FromSender string    `orm:"column(from_sender)"`
	Type       string    `orm:"column(type)"`
	
}

// TableName specifies the table name in the database
func (t *SMSTemplate) TableName() string {
	return "ls_sms_template"
}

// func init() {
// 	// Register model with Beego ORM
// 	orm.RegisterModel(new(LSSMSTemplate))
// }
