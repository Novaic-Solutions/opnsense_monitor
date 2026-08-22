package client

import (
	"bytes"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"context"
	"sync"
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

//----------------------------------------------------------------------------
// Create a response object from the HTTP response received from the API call.
// Use the ResponseObjType field in the ApiRequeest object to determine
// which type of object to create from the response data. The response data is
// unmarshaled into the appropriate object type and set in the Data field of the
// EndpointResponse object.
//----------------------------------------------------------------------------
func (apiReq *FirewallLogStreamClient) CreateResponseObj(httpResp *http.Response) config.EndpointResponse {
	response := config.EndpointResponse{
		Uri: apiReq.ApiRequest.Uri,
		Timestamp: time.Now().Format(time.RFC3339),
		ResponseDataType: apiReq.ApiRequest.ResponseObjType,
		Data: nil,
	}

	byteArr, err := GetResponseData(httpResp)
	if err != nil {
		fmt.Printf("Firewall_Log_Stream.go -- Error getting response data: %v\n", err)
		return response
	}

	//---------------------------------------------------------------------------
	// Unmarshal the response data into the appropriate object type based on the
	// ResponseObjType field in the ApiRequest object.
	//---------------------------------------------------------------------------
	switch apiReq.ApiRequest.ResponseObjType {
	case "FirewallLogEntry":
		var logEntries []data.FirewallLogEntry

		if err := json.Unmarshal(byteArr, &logEntries); err != nil {
			fmt.Printf("Firewall_Log_Stream.go -- Error unmarshaling response body to []FirewallLogEntry: %v\n", err)
		} else {
			response.Data = logEntries
		}
	}

	return response
}

//----------------------------------------------------------------------------
// Perform the HTTP call
//----------------------------------------------------------------------------
func (apiReq *FirewallLogStreamClient) Call(ctx context.Context, wg *sync.WaitGroup) {
	// This might cause an issue with the wait group since this function is recursive. Need to figure out how to handle this.
	defer wg.Done()
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
		req, err = http.NewRequest(apiReq.ApiRequest.Method, url, nil)
	} else {
		//---------------------------------------------------------------------
		// Ensure only the content type for application/json is set IF
		// the request body is not empty. This prevents 400 errors from the API.
		//---------------------------------------------------------------------
		req, err = http.NewRequest(apiReq.ApiRequest.Method, url, bytes.NewBuffer(apiReq.ApiRequest.Body))
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
	// Send the request
	//---------------------------------------------------------------------------
	response, _ := SendRequest(req)

	apiResponse := apiReq.CreateResponseObj(response)

	//---------------------------------------------------------------------------
	// Send the response through the channel to the server for processing and storage.
	//---------------------------------------------------------------------------
	apiReq.ResponseChannel <- apiResponse

	time.Sleep(15 * time.Second)

	apiReq.ApiRequest.Params["digest"] = apiResponse.Data.([]data.FirewallLogEntry)[len(apiResponse.Data.([]data.FirewallLogEntry))-1].Digest
	apiReq.ApiRequest.Params["limit"] = "10000"
	apiReq.Call(ctx, wg)
}
