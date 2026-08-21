package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"time"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"github.com/Novaic-Solutions/opnsense_monitor/data"
)

type Request struct {
	ApiRequest *config.ApiRequest
	ResponseChannel chan config.EndpointResponse
}

//----------------------------------------------------------------------------
// Create new client
//----------------------------------------------------------------------------
func NewRequest(apiRequest *config.ApiRequest, responseChannel chan config.EndpointResponse) *Request {
	return &Request{
		ApiRequest: apiRequest,
		ResponseChannel: responseChannel,
	}
}

//----------------------------------------------------------------------------
// Create response object based on the ApiRequest.ResponseObjType field.
//----------------------------------------------------------------------------
func (req *Request) CreateResponseObj(httpResp *http.Response) config.EndpointResponse {
	response := config.EndpointResponse{
		Uri: req.ApiRequest.Uri,
		Timestamp: time.Now().Format(time.RFC3339),
		ResponseDataType: req.ApiRequest.ResponseObjType,
		Data: nil,
	}

	byteArr, err := GetResponseData(httpResp)
	if err != nil {
		fmt.Printf("Request.go -- Error getting response data: %v\n", err)
		return response
	}

	switch req.ApiRequest.ResponseObjType {
	case "ArpTable":
		var arpTable = data.ArpTable{}
		if err := json.Unmarshal(byteArr, &arpTable); err != nil {
			fmt.Printf("Request.go -- Error unmarshaling response body to ArpTable: %v\n", err)
		} else {
			response.Data = arpTable
		}
	case "IfaceStatistics":
		var stats = data.IfaceStatistics{}
		if err := json.Unmarshal(byteArr, &stats); err != nil {
			fmt.Printf("Request.go -- Error unmarshaling response body to IfaceStatistics: %v\n", err)
		} else {
			response.Data = stats
		}
	case "FirewallSessions":
		var sessions data.FirewallSessions
		if err := json.Unmarshal(byteArr, &sessions); err != nil {
			fmt.Printf("Request.go -- Error unmarshaling response body to FirewallSessions: %v\n", err)
		} else {
			response.Data = sessions
		}
	case "FirewallStates":
		var states = data.FirewallStates{}
		if err := json.Unmarshal(byteArr, &states); err != nil {
			fmt.Printf("Request.go -- Error unmarshaling response body to FirewallStates: %v\n", err)
		} else {
			response.Data = states
		}
	case "Interfaces":
		var interfaces = data.Interfaces{}
		if err := json.Unmarshal(byteArr, &interfaces); err != nil {
			fmt.Printf("Request.go -- Error unmarshaling response body to Interfaces: %v\n", err)
		} else {
			response.Data = interfaces
		}
	default:
		fmt.Printf("Request.go -- Unknown ResponseObjType: %s\n", req.ApiRequest.ResponseObjType)
	}

	return response
}

//----------------------------------------------------------------------------
// 
//----------------------------------------------------------------------------
func (req *Request) Call() {
	var httpReq *http.Request
	var err error

	url := req.ApiRequest.Url + req.ApiRequest.Uri

	//---------------------------------------------------------------------------
	// If there are any parameters, append them to the URL as a query string.
	//---------------------------------------------------------------------------
	if len(req.ApiRequest.Params) > 0 {
		url += "?"
		for key, value := range req.ApiRequest.Params {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1] // Remove the trailing '&'know
	}

	//---------------------------------------------------------------------------
	// Create a new request object with the appropriate method, endpoint, and body.
	// If the body is empty, set it to nil to prevent sending an empty JSON object.
	// This should help prevent 400 errors from the API.
	//---------------------------------------------------------------------------
	if len(req.ApiRequest.Body) == 0 {
		httpReq, err = http.NewRequest(req.ApiRequest.Method, url, nil)
		if err != nil {
			fmt.Printf("Request.go -- Error creating HTTP request: %v\n", err)
			return
		}
	} else {
		//---------------------------------------------------------------------
		// Ensure only the content type for application/json is set IF
		// the request body is not empty. This prevents 400 errors from the API.
		//---------------------------------------------------------------------
		httpReq, err = http.NewRequest(req.ApiRequest.Method, url, bytes.NewBuffer(req.ApiRequest.Body))
		if err != nil {
			fmt.Printf("Request.go -- Error creating HTTP request: %v\n", err)
			return
		}
		httpReq.Header.Set("Content-Type", "application/json")
	}

	//---------------------------------------------------------------------------
	// Set the basic auth for the request using
	// the token created in opnsense.
	//---------------------------------------------------------------------------
	httpReq.SetBasicAuth(req.ApiRequest.Username, req.ApiRequest.Password)

	//---------------------------------------------------------------------------
	// Send the request
	//---------------------------------------------------------------------------
	response, _ := SendRequest(httpReq)

	apiResponse := req.CreateResponseObj(response)

	//---------------------------------------------------------------------------
	// Send the response through the channel to the server for processing and storage.
	//---------------------------------------------------------------------------
	req.ResponseChannel <- apiResponse

	time.Sleep(15 * time.Second)
	req.Call()
}
