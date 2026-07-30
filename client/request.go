package client

import (
	"fmt"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"time"
	"io"
	"encoding/json"
)

type Request struct {
	ApiRequest *ApiRequest
	ResponseChannel chan EndpointResponse
}

//----------------------------------------------------------------------------
// Create new client
//----------------------------------------------------------------------------
func NewRequest(apiRequest *ApiRequest, responseChannel chan EndpointResponse) *Request {
	return &Request{
		ApiRequest: apiRequest,
		ResponseChannel: responseChannel,
	}
}

//----------------------------------------------------------------------------
// Client Gather function
// Send request to API endpoint and gather response, then place on the 
// channel for processing by the server.
//----------------------------------------------------------------------------
func (req *Request) Gather() {
	resp, err := req.ApiRequest.SendRequest()

	if err != nil {
		fmt.Printf("Client.go -- Error sending request: %v", err)
		return 
	}
	defer resp.Body.Close()


	//----------------------------------------------------------------------------
	// Read the response body
	//----------------------------------------------------------------------------
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Client.go -- Error reading response body: %v", err)
		return
	}


	//----------------------------------------------------------------------------
	// Create dataResult map to hold the response data
	// in a format that can hold the data from the body after it is 
	// serialized from JSON.
	//----------------------------------------------------------------------------
	var dataResult map[string]interface{}
	if err := json.Unmarshal(body, &dataResult); err != nil {
		fmt.Println("---------------------------------------------------")
		fmt.Printf("Client.go -- Error unmarshaling response body: %v\n", err)
		fmt.Printf("Client.go -- URI: %s\n", req.ApiRequest.Endpoint.Request.Uri)
		fmt.Printf("Client.go -- Response Body: %s\n", string(body))
		fmt.Println("---------------------------------------------------")
		return
	}

	// Add logic here to process the response and send it through the channel
	// for further processing

	//---------------------------------------------------------------------
	// Put the response data on the channel for processing 
	// by the server.
	//---------------------------------------------------------------------
	req.ResponseChannel <- EndpointResponse{
		Uri: req.ApiRequest.Endpoint.Request.Uri,
		Timestamp: time.Now().Format(time.RFC3339),
		Data: dataResult,
	}

}
