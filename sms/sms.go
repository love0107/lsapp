package sms

type Request struct{
	Mobile string `json:"mobile"`
	OTP string `json:"otp"`
}
type Response struct{
	Status string
	StatusCode int
	Body string
}