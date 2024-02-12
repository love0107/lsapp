package clickatell

import "lsapp/sms"

// base_url := "https://platform.clickatell.com/messages/http/send"
// apiKey := "sQ5Nao4iTJujnhM1J3MS9Q=="
// Kicktail is a struct representing the Kicktail SMS vendor.
type Cicktail struct {
	// Cicktail-specific configuration or credentials
}

// SendSMS implements the SMSVendor interface for Kicktail.
func (c Cicktail) SendSMS(request sms.Request) (resp sms.Response,err error) {
	
	return resp, nil
}
