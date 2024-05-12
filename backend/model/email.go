package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// LSEmailTemplate represents the ls_email_template table in the database
type EmailTemplate struct {
	Id        int       `orm:"column(id);auto"`
	Name      string    `orm:"column(name)"`
	Subject   string    `orm:"column(subject)"`
	Body      string    `orm:"column(body)"`
	Bcc       string    `orm:"column(bcc);null"`
	Cc        string    `orm:"column(cc);null"`
	To        string    `orm:"column(to);null"`
	FromName  string    `orm:"column(fromName)"`
	FromEmail string    `orm:"column(fromEmail)"`
	ReplyTo   string    `orm:"column(reply_to);null"`
	Type      string    `orm:"column(type)"`
	CreatedOn time.Time `orm:"column(created_on);type(timestamp);auto_now_add"`
	UpdatedOn time.Time `orm:"column(updated_on);type(timestamp);auto_now"`
}

// TableName specifies the table name in the database
func (e *EmailTemplate) TableName() string {
	return "ls_email_template"
}

func (e *EmailTemplate) GetEmailTemplateByName(name string) (*EmailTemplate, error) {
	o := orm.NewOrm()
	template := &EmailTemplate{}
	err := o.QueryTable(new(EmailTemplate)).Filter("name", name).One(template)
	if err != nil {
		return nil, err
	}
	return template, nil
}
