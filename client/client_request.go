package client

import (
	"crypto/tls"
	"fmt"
	"io"
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

	return resp, nil
}

//----------------------------------------------------------------------------
//	Get response body data from the API endpoint and return it as a byte slice.
//----------------------------------------------------------------------------
func GetResponseData(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, fmt.Errorf("response is nil")
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v", err)
		return nil, err
	}

	return body, nil
}