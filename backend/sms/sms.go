package sms

import (
	"errors"
	"log"
	"lsapp/model"
	"lsapp/sms/clickatell"
)

const (
	constClickatell = "clickatell"
)

type SmsService interface {
	SendSms(clickatell.Request) (clickatell.Response, error)
}

var Provider map[string]SmsService

type Request struct {
	To      string
	Type    string
	Message string
	Vendor  string
	UserId  int64
	Config  map[string]string
}
type Response struct {
	Code   int
	Status string
	Body   string
}

func init() {
	Provider = make(map[string]SmsService)
	Provider[constClickatell] = clickatell.Clickatell{Name: constClickatell}
}

// here need to create the service layer and after that we call send function
func SendSms(request Request) (response Response, err error) {
	log.Println("Request ::", request)
	if Provider[request.Vendor] == nil {
		log.Println("no provider found")
		return response, errors.New("failed to make object")
	}
	configMap, err := model.GetConfigByType(constClickatell)
	if err != nil {
		log.Println("failed to get config:", constClickatell, " error:", err)
		return response, errors.New("failed to get configs")
	}
	request.Config = configMap
	resp, err := Provider[request.Vendor].SendSms(clickatell.Request(request))
	if err != nil {
		log.Println("failed to send sms for request", request)
	}
	response = Response(resp)

	// model.AddSmsLog(response)

	return response, nil

}
