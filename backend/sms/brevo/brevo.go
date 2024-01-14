package brevo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"lsapp/model"
	"lsapp/sms"
	"net/http"
)

func sendSMS(request map[string]interface{}) (response sms.Respones, err error) {
	functionName := "brevo.sendSMS"
	config, err := model.GetConfigByType("brevo/url")
	url := config["brevo/url"]
	apiKey := config["brevo/apikey"]
	payloadBytes, err := json.Marshal(request)
	if err != nil {
		fmt.Println(functionName, " error Marshal Request ", err)
		return response, errors.New("marshal error")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", apiKey)

	// Create a new HTTP client
	client := &http.Client{}

	// Send request and get response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()
	return
}
