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

// Create Request with Basic Auth
func CreateRequest(method, url, username, password string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)
	return req, nil
}

// Send Request and return Response
func SendRequest(client *http.Client, req *http.Request) (*http.Response, error)