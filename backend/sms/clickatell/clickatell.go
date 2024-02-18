package clickatell

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Kicktail is a struct representing the Kicktail SMS vendor.
type Clickatell struct {
	Name string
}
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

// SendSMS implements the SMSVendor interface for Kicktail.
func (c Clickatell) SendSms(request Request) (response Response, err error) {

	// Create a new GET request
	url, err := getFullUrl(request)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	// Send the request using the default HTTP client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}
	response.Body = string(body)
	response.Code = resp.StatusCode
	response.Status = resp.Status
	return response, nil
}
// get the full url
func getFullUrl(request Request) (fullUrl string, err error) {
	baseUrl := request.Config["url"]
	apiKey := request.Config["apiKey"]
	if baseUrl == "" || apiKey == "" {
		log.Println("empty base url or apiKey")
		return "", errors.New("empty apiKey or url")
	}
	url := url.Values{}
	url.Add("apiKey", apiKey)
	url.Add("to", request.To)
	url.Add("content", request.Message)
	fullUrl = baseUrl + url.Encode()
	return fullUrl, nil
}
