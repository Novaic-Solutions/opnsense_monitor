package client

import (
	"net/http"
	"gitlab.com/tymonx/go-formatter/formatter"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
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


func (req *Request) Call() {

}