package client

// import (
// 	"fmt"
// 	"time"
// 	"io"
// 	"encoding/json"
// 	"github.com/Novaic-Solutions/opnsense_monitor/config"
// 	"gitlab.com/tymonx/go-formatter/formatter"
// )

type FirewallLogStreamClient struct {
	ApiRequest *ApiRequest
	ResponseChannel chan EndpointResponse
}

func NewFirewallLogStreamClient(apiRequest *ApiRequest, responseChannel chan EndpointResponse) *FirewallLogStreamClient {
	return &FirewallLogStreamClient{
		ApiRequest: apiRequest,
		ResponseChannel: responseChannel,
	}
}


func (req *FirewallLogStreamClient) Call() {
	
}
