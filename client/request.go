package client

import (
	"net/http"
	"gitlab.com/tymonx/go-formatter/formatter"
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



func (req *Request) Call() {
	var req *http.Request
	var err error

	url := req.ApiRequest.Endpoint.Request.Urir

	if len(req.ApiRequest.Endpoint.Request.Params) > 0 {
		formattedParams, _ := formatter.Format(req.ApiRequest.Endpoint.Request.Params)
		// 	formattedParams, _ := formatter.Format(cli.ApiRequest.Request.Params )
		// initialString := "https://1.1.1.1:666/api/diagnostics/firewall/log/?digest={digest}&limit={limit}"
		// digest := ""
		// formattedInitString, _ := formatter.Format(initialString, formatter.Named{
		// 	"digest": digest,
		// 	"limit":  "10000",
		// })
	
}