package util

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func MakeRequest(url string, payload string, method string, headers map[string]string) (string, error) {
	// Create a request with the specified method and the provided URL
	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		return "", err
	}

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Create an HTTP client
	client := &http.Client{}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check the response status code (you may customize this based on your API)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	return string(responseBody), nil
}

