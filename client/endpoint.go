package client

import (
	"crypto/tls"
	"net/http"
	"bytes"
	"time"
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
	req, err := http.NewRequest(apiReq.Method, apiReq.Endpoint, bytes.NewBuffer(apiReq.Body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(apiReq.Username, apiReq.Password)

	client := CreateHTTPClient(30)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
