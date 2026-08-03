package client

import (
	"bytes"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"github.com/Novaic-Solutions/opnsense_monitor/data"
)

type FirewallLogStreamClient struct {
	ApiRequest *config.ApiRequest
	ResponseChannel chan config.EndpointResponse
}

func NewFirewallLogStreamClient(apiRequest *config.ApiRequest, responseChannel chan config.EndpointResponse) *FirewallLogStreamClient {
	return &FirewallLogStreamClient{
		ApiRequest: apiRequest,
		ResponseChannel: responseChannel,
	}
}

func (apiReq *FirewallLogStreamClient) CreateResponseObj(httpResp *http.Response) config.EndpointResponse {
	response := config.EndpointResponse{
		Uri: apiReq.ApiRequest.Uri,
		Timestamp: time.Now().Format(time.RFC3339),
		ResponseDataType: apiReq.ApiRequest.ResponseType,
		Data: nil,
	}

	byteArr, err := GetResponseData(httpResp)
	if err != nil {
		fmt.Printf("Error getting response data: %v\n", err)
		return response
	}
	
	switch apiReq.ApiRequest.ResponseObjType {
	case "FirewallLogEntry":
		var logEntry data.FirewallLogEntry
		if err := json.Unmarshal(byteArr, &logEntry); err != nil {
			fmt.Printf("Error unmarshaling response body to FirewallLogEntry: %v\n", err)
		}
		response.Data = logEntry
	}


	return response
}


func (apiReq *FirewallLogStreamClient) Call() {
	var req *http.Request
	var err error

	url := apiReq.ApiRequest.Url + apiReq.ApiRequest.Uri

	//---------------------------------------------------------------------------
	// If there are any parameters, append them to the URL as a query string.
	//---------------------------------------------------------------------------
	if len(apiReq.ApiRequest.Params) > 0 {
		url += "?"
		for key, value := range apiReq.ApiRequest.Params {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1] // Remove the trailing '&'
	}


	//---------------------------------------------------------------------------
	// Create a new request object with the appropriate method, endpoint, and body.
	// If the body is empty, set it to nil to prevent sending an empty JSON object.
	// This should help prevent 400 errors from the API.
	//---------------------------------------------------------------------------
	if len(apiReq.ApiRequest.Body) == 0 {
		req, err = http.NewRequest(apiReq.ApiRequest.Method, apiReq.ApiRequest.Url+apiReq.ApiRequest.Uri, nil)
	} else {
		//---------------------------------------------------------------------
		// Ensure only the content type for application/json is set IF
		// the request body is not empty. This prevents 400 errors from the API.
		//---------------------------------------------------------------------
		req, err = http.NewRequest(apiReq.ApiRequest.Method, apiReq.ApiRequest.Url+apiReq.ApiRequest.Uri, bytes.NewBuffer(apiReq.ApiRequest.Body))
		req.Header.Set("Content-Type", "application/json")
	}
	
	if err != nil {
		return
	}

	//---------------------------------------------------------------------------
	// Set the basic auth for the request using
	// the token created in opnsense.
	//---------------------------------------------------------------------------
	req.SetBasicAuth(apiReq.ApiRequest.Username, apiReq.ApiRequest.Password)
	
	//---------------------------------------------------------------------------
	// Set the basic auth for the request using
	// the token created in opnsense.
	//---------------------------------------------------------------------------
	req.SetBasicAuth(apiReq.ApiRequest.Username, apiReq.ApiRequest.Password)

	//---------------------------------------------------------------------------
	// Send the request
	//---------------------------------------------------------------------------
	response, _ := SendRequest(req)

	apiResponse := apiReq.CreateResponseObj(response)

	//---------------------------------------------------------------------------
	// Send the response through the channel to the server for processing and storage.
	//---------------------------------------------------------------------------
	apiReq.ResponseChannel <- apiResponse


}
