package client

import (
	"fmt"
)




//----------------------------------------------------------------------------
//  Used to store the response from the API endpoint and send it through
//  a channel to the server for processing and storage in the database.
//----------------------------------------------------------------------------


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
func PopulateApiRequests(endPoints []Endpoint, responseChannel chan EndpointResponse, usr, pwd string) ([]Caller, error) {
	ClientCallers := make([]Caller, 0, len(endPoints))

	for _, endP := range endPoints {
		apiReq := ApiRequest{
			Endpoint: endP,
			Username: usr,
			Password: pwd,
		}
		client := NewCaller(endP.Type, &apiReq, responseChannel)
		ClientCallers = append(ClientCallers, client)
	}
	return ClientCallers, nil
}