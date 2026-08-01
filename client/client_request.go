package client

import (
	//"bytes"
	"crypto/tls"
	//"fmt"
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

// // ------------------------------------------------------------------------------
// // Create Request with Basic Auth and send it
// // ------------------------------------------------------------------------------
// func (apiReq *ApiRequest) SendRequest() (*http.Response, error) {
// 	var req *http.Request
// 	var err error

// 	//---------------------------------------------------------------------------
// 	// Create a new request object with the appropriate method, endpoint, and body.
// 	// If the body is empty, set it to nil to prevent sending an empty JSON object.
// 	// This should help prevent 400 errors from the API.
// 	//---------------------------------------------------------------------------
// 	if len(apiReq.Body) == 0 {
// 		req, err = http.NewRequest(apiReq.Method, apiReq.Endpoint, nil)
// 	} else {
// 		//---------------------------------------------------------------------
// 		// Ensure only the content type for application/json is set IF
// 		// the request body is not empty. This prevents 400 errors from the API.
// 		//---------------------------------------------------------------------
// 		req, err = http.NewRequest(apiReq.Method, apiReq.Endpoint, bytes.NewBuffer(apiReq.Body))
// 		req.Header.Set("Content-Type", "application/json")
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	//---------------------------------------------------------------------------
// 	// Set the basic auth for the request using
// 	// the token created in opnsense.
// 	//---------------------------------------------------------------------------
// 	req.SetBasicAuth(apiReq.Username, apiReq.Password)
	
// 	//---------------------------------------------------------------------------
// 	// Set the basic auth for the request using
// 	// the token created in opnsense.
// 	//---------------------------------------------------------------------------
// 	req.SetBasicAuth(apiReq.Username, apiReq.Password)

// 	//----------------------------------------------------------------------------
// 	// Create the HTTP client
// 	//----------------------------------------------------------------------------
// 	client := CreateHTTPClient(30)

// 	//----------------------------------------------------------------------------
// 	// Send the request and return the response
// 	//----------------------------------------------------------------------------
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Printf("Endpoint.go -- Request sent to %s with method %s\n", apiReq.Endpoint, apiReq.Method)
// 	fmt.Printf("Endpoint.go -- Response Status: %s\n", resp.Status)

// 	// Check if the response body is empty and print it for debugging purposes.
// 	if apiReq.Body != nil {
// 		fmt.Printf("Endpoint.go -- Request Body: %s\n", string(apiReq.Body))
// 	} else {
// 		fmt.Printf("Endpoint.go -- Request Body: nil\n")
// 	}

// 	return resp, nil
// }
