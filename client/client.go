package client

// import (
// 	"github.com/Novaic-Solutions/opnsense_monitor/config"
// 	"fmt"
// 	"encoding/json"
// 	"io"
// 	"time"
// )

//----------------------------------------------------------------------
// Endpoint represents a single API Endpoint to call and the 
// data required to perform the request.
//----------------------------------------------------------------------
type Endpoint struct {
	Request RequestObj `yaml:"request"`
	Type string `yaml:"type"`
	ResponseObjType string `yaml:"response_obj_type"`
}

type RequestObj struct {
	Uri string `yaml:"uri"`
	Method string `yaml:"method"`
	ResponseType string `yaml:"response_type"`
	Params string `yaml:"params"`
	RequestBody any `yaml:"request_body"`
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
	Endpoint Endpoint
	Username string
	Password string
}

type Caller interface {
	Call()
}

//----------------------------------------------------------------------------
//  Create a new client for the given endpoint and return it.
//----------------------------------------------------------------------------
func NewCaller(reqType string, apiRequest *ApiRequest, responseChannel chan EndpointResponse) Caller {
	switch reqType {
	case "request":
		return NewRequest(apiRequest, responseChannel)
	case "firewall_log_stream":
		return NewFirewallLogStreamClient(apiRequest, responseChannel)
	default:
		fmt.Printf("Client.go -- Unknown client type: %s\n", reqType)
		return nil
	}
}

//---------------------------------------------------------------------------
//  Populate the slice of clients with a client 
//  for each endpoint in the config file.
//---------------------------------------------------------------------------
func PopulateApiRequests(conf *config.Config) ([]*Caller, error) {




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