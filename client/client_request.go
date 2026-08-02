package client

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

// ------------------------------------------------------------------------------
// Create TLS Config
// ------------------------------------------------------------------------------
func CreateTlsConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}

// ------------------------------------------------------------------------------
// Create Transport for HTTP Client
// ------------------------------------------------------------------------------
func CreateHTTPTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: CreateTlsConfig(),
	}
}

// ------------------------------------------------------------------------------
// Create HTTP Client
// ------------------------------------------------------------------------------
func CreateHTTPClient(timeout int) *http.Client {
	return &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: CreateHTTPTransport(),
	}
}

// ------------------------------------------------------------------------------
// Create Request with Basic Auth and send it
// ------------------------------------------------------------------------------
func SendRequest(apiReq *http.Request) (*http.Response, error) {

	//----------------------------------------------------------------------------
	// Create the HTTP client
	//----------------------------------------------------------------------------
	client := CreateHTTPClient(30)

	//----------------------------------------------------------------------------
	// Send the request and return the response
	//----------------------------------------------------------------------------
	resp, err := client.Do(apiReq)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Endpoint.go -- Request sent to %s with method %s\n", apiReq.URL, apiReq.Method)
	fmt.Printf("Endpoint.go -- Response Status: %s\n", resp.Status)

	// Check if the response body is empty and print it for debugging purposes.
	if apiReq.Body != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(apiReq.Body)
		fmt.Printf("Endpoint.go -- Request Body: %s\n", buf.String())
	} else {
		fmt.Printf("Endpoint.go -- Request Body: nil\n")
	}

	return resp, nil
}
