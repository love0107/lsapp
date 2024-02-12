package sms

import (
	"encoding/json"
	"lsapp/log"
	"lsapp/util"
	"net/http"
)


// SMSVendor is an interface for sending SMS.
type SMSVendor interface {
	SendSMS(request Request) (Response, error)
 }

type Request struct {
	Mobile string `json:"mobile"`
	OTP    string `json:"otp"`
}
type Response struct {
	Status     string
	StatusCode int
	Body       string
}

// SendSMS in the gateway layer uses the SMSVendor interface for abstraction.
func SendSMS(request Request, vendor SMSVendor) (Response, error) {
	// Perform any common logic

	// Use the provided vendor to send the SMS
	response, err := vendor.SendSMS(request)
	if err != nil {
		// Handle errors
		return Response{}, err
	}

	// Perform any post-processing logic

	return response, nil
}



func SendSms(request Request) (resp Response, err error) {
	functionName := "<SendSms>"
	log.Println(functionName, " request:", request)

	byteRequestBody, err := json.Marshal(request)
	if err != nil {
		log.Println(functionName, " failed to marshal request body. Error:", err)
		return resp, err
	}

	// Initialize headers map with key and values if needed
	headers := map[string]string{}

	url := "<your_api_url>"
	apiResponse, err := util.MakeRequest(url, string(byteRequestBody), http.MethodPost, headers)
	if err != nil {
		log.Println(functionName, " failed to make request. Error:", err)
		return resp, err
	}

	err = json.Unmarshal([]byte(apiResponse), &resp)
	if err != nil {
		log.Println(functionName, " failed to unmarshal API response. Error:", err)
		return resp, err
	}

	return resp, nil
}
