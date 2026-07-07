package client

import (
	"crypto/tls"
	"net/http"
	"bytes"
	"time"
	"fmt"
)

type ApiRequest struct {
	Method string
	Endpoint string
	ResponseType string
	Body []byte
	Username string
	Password string
}

//------------------------------------------------------------------------------
// Create TLS Config
//------------------------------------------------------------------------------

func CreateTlsConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}

//------------------------------------------------------------------------------
// Create Transport for HTTP Client
//------------------------------------------------------------------------------
func CreateHTTPTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: CreateTlsConfig(),
	}
}

//------------------------------------------------------------------------------
// Create HTTP Client
//------------------------------------------------------------------------------
func CreateHTTPClient(timeout int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: CreateHTTPTransport(),
	}
}

//------------------------------------------------------------------------------
// Create Request with Basic Auth and send it
//------------------------------------------------------------------------------
func (apiReq *ApiRequest) SendRequest() (*http.Response, error) {
	var req *http.Request
	var err error

	if len(apiReq.Body) == 0 {
		req, err = http.NewRequest(apiReq.Method, apiReq.Endpoint, nil)
	} else {
		//---------------------------------------------------------------------
		// Ensure only the content type for application/json is set IF 
		// the request body is not empty. This prevents 400 errors from the API.
		//---------------------------------------------------------------------
		req, err = http.NewRequest(apiReq.Method, apiReq.Endpoint, bytes.NewBuffer(apiReq.Body))
		req.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(apiReq.Username, apiReq.Password)

	client := CreateHTTPClient(30)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Request sent to %s with method %s\n", apiReq.Endpoint, apiReq.Method)
	fmt.Printf("Response Status: %s\n", resp.Status)
	if apiReq.Body != nil {
		fmt.Printf("Request Body: %s\n", string(apiReq.Body))
	} else {
		fmt.Printf("Request Body: nil\n")
	}

	return resp, nil
}
