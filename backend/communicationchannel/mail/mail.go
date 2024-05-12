package mail

import (
	"log"
	"lsapp/communicationchannel/mail/smtpmail"
	"lsapp/model"
	"lsapp/util"
)

const (
	constSMTP = "smtp"
)

var Provider map[string]MailService



type MailService interface {
	SendMail(util.Request) (util.Response, error)
}

func init() {
	Provider = make(map[string]MailService)
	Provider[constSMTP] = smtpmail.SMTP{Name: constSMTP}
}

// mail service send mail
func SendMail(request util.Request) (response util.Response, err error) {
	functionName := "<SendMail> "
	template, err := new(model.EmailTemplate).GetEmailTemplateByName(request.TemplateName)
	if err != nil {
		log.Println(functionName, "failed to get template:", request.TemplateName, " error:", err)
		return response, err
	}
	return
}
