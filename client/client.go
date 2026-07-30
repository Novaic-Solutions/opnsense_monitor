package client

import (
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"fmt"
	"encoding/json"
	"io"
	"time"
)

//----------------------------------------------------------------------
// Endpoint represents a single API Endpoint to call and the 
// data required to perform the request.
//----------------------------------------------------------------------
type RequestObj struct {
	Uri string `yaml:"uri"`
	Method string `yaml:"method"`
	ResponseType string `yaml:"response_type"`
	Params string `yaml:"params"`
	RequestBody any `yaml:"request_body"`
}

type Endpoint struct {
	Request RequestObj `yaml:"request"`
	Type string `yaml:"type"`
	ResponseObjType string `yaml:"response_obj_type"`
}

//----------------------------------------------------------------------------
//  Used to store the response from the API endpoint and send it through
//  a channel to the server for processing and storage in the database.
//----------------------------------------------------------------------------
type EndpointResponse struct {
	Uri string
	Timestamp string
	Data map[string]interface{}
}

type ApiRequest struct {
	Request Endpoint
	Username string
	Password string
}

type Client struct {
	ApiRequest *ApiRequest
	ResponseChannel chan EndpointResponse
}

// 	TODO: Use an interface for the client so that the server can use either a Client or a StreamClient,
//  depending on the type of endpoint being monitored.
type Caller interface {
	Gather()
}
    
//----------------------------------------------------------------------------
// Populate API Requests from Config
//----------------------------------------------------------------------------
// func PopulateApiRequests(conf *config.Config) (*[]ApiRequest, error) {
	
// 	//-------------------------------------------------------------------------------
// 	// Create a slice of ApiRequest objects to hold the requests to be sent to the API endpoints.
// 	//-------------------------------------------------------------------------------
// 	apiObj := make([]ApiRequest, 0, 100)

// 	//-------------------------------------------------------------------------------
// 	// Loop through the endpoints in the config and create an ApiRequest object for each one.
// 	//-------------------------------------------------------------------------------
// 	for _, value := range conf.API.Endpoints {

// 		endp := value.Request.RequestBody.(map[string]interface{})

// 		bytes, err := json.Marshal(endp)
// 		if err != nil {
// 			fmt.Printf("Client.go -- Error marshaling to JSON: %v", err)
// 		}

// 		//-------------------------------------------------------------------------------------
// 		// Prevents the body being sent as an empty JSON object when it is not needed for the request.
// 		// This should help prevent 400 errors.
// 		//-------------------------------------------------------------------------------------
// 		if len(endp) == 0 {
// 			bytes = nil
// 		}

// 		newReq := ApiRequest{
// 			Method: value.Request.Method,
// 			Endpoint: conf.API.BaseURL + value.Request.Uri,
// 			ResponseType: value.Request.ResponseType,
// 			Username: conf.API.ApiKey,
// 			Password: conf.API.ApiSecret,
// 			Body: bytes,
// 		}
// 		apiObj = append(apiObj, newReq)
// 	}

// 	if len(apiObj) == 0 {
// 		return nil, fmt.Errorf("Client.go -- No API requests found in configuration")
// 	}

// 	return &apiObj, nil
// }

//----------------------------------------------------------------------------
// Create new client
//----------------------------------------------------------------------------
func NewClient(apiRequest *ApiRequest, responseChannel chan EndpointResponse) *Client {
	return &Client{
		ApiRequest: apiRequest,
		ResponseChannel: responseChannel,
	}
}

//----------------------------------------------------------------------------
// Client Gather function
// Send request to API endpoint and gather response, then place on the 
// channel for processing by the server.
//----------------------------------------------------------------------------
func (cli *Client) Gather() {
	resp, err := cli.ApiRequest.SendRequest()

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
		fmt.Printf("Client.go -- URI: %s\n", cli.ApiRequest.Endpoint)
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
	cli.ResponseChannel <- EndpointResponse{
		Uri: cli.ApiRequest.Endpoint,
		Timestamp: time.Now().Format(time.RFC3339),
		Data: dataResult,
	}

}
