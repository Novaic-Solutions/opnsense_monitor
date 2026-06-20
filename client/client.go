package client

import (
	"crypto/tls"
	"net/http"
	"bytes"
	"time"
)


// Create TLS Config
func CreateTlsConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}

// Create Transport for HTTP Client
func CreateHTTPTransport() *http.Transport{
	return &http.Transport{
		TLSClientConfig: CreateTlsConfig(),
	}
}

// Create HTTP Client
func CreateHTTPClient(timeout int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: CreateHTTPTransport(),
	}
}

// Create Request with Basic Auth and send it
func SendRequest(method, url, username, password string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	client := CreateHTTPClient(30)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
