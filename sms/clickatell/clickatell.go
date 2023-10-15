package clickatell

import (
	"fmt"
	"io"
	"lsapp/model"
	"lsapp/sms"
	"net/http"
	"net/url"
)

// apiKey := "sQ5Nao4iTJujnhM1J3MS9Q=="
const (
	SMS_CLICKATELL = "sms/clickatell"
	URL            = "url"
	API_KEY        = "apikey"
)

// send the otp
func SendOTP(req sms.Request) {
	// url := "https://platform.clickatell.com/messages/http/send"
	// get the config from the db
	configMap, err := model.GetConfigByType(SMS_CLICKATELL)
	if err != nil {
		fmt.Println("error while fetching the ls config: ", err)
		return
	}
	_apiKey := configMap[API_KEY]
	baseurl := configMap[URL]
	to := req.Mobile
	content := req.OTP

	// prepare the url
	queryParams := url.Values{}
	queryParams.Set("apiKey", _apiKey)
	queryParams.Set("to", to)
	queryParams.Set("content", content)
	url := baseurl + "?" + queryParams.Encode()
	// Send GET request
	resp, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Error: fail to send the otp: ", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("body", string(body), " and resp---->> ", resp, "resp stats response-->", resp.Response)
}
