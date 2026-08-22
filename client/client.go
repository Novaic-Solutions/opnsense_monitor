package client

import (
	"context"
	"fmt"
	"sync"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
)


//----------------------------------------------------------------------------
//  Used to store the response from the API endpoint and send it through
//  a channel to the server for processing and storage in the database.
//----------------------------------------------------------------------------


type Caller interface {
	Call(ctx context.Context, wg *sync.WaitGroup)
}

//----------------------------------------------------------------------------
//  Create a new client for the given endpoint and return it.
//----------------------------------------------------------------------------
func NewCaller(apiRequest *config.ApiRequest, responseChannel chan config.EndpointResponse) Caller {
	switch apiRequest.TypeRequest {
	case "request":
		return NewRequest(apiRequest, responseChannel)
	case "firewall_log_stream":
		return NewFirewallLogStreamClient(apiRequest, responseChannel)
	default:
		fmt.Printf("Client.go -- Unknown client type: %s\n", apiRequest.TypeRequest)
		return nil
	}
}

