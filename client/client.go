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

