package client

// import (
// 	"fmt"
// 	"time"
// 	"io"
// 	"encoding/json"
// 	"github.com/Novaic-Solutions/opnsense_monitor/config"
// 	"gitlab.com/tymonx/go-formatter/formatter"
// )

// type FirewallLogStreamClient struct {
// 	ApiRequest *ApiRequest
// 	ResponseChannel chan EndpointResponse
// }

// func NewFirewallLogStreamClient(apiRequest *ApiRequest, responseChannel chan EndpointResponse) *FirewallLogStreamClient {
// 	return &FirewallLogStreamClient{
// 		ApiRequest: apiRequest,
// 		ResponseChannel: responseChannel,
// 	}
// }


// //----------------------------------------------------------------------------
// // Stream Client Gather function
// // Send request to API endpoint and gather response, then place on the
// //----------------------------------------------------------------------------
// func (cli *FirewallLogStreamClient) Gather(token,limit string) {

// 	formattedParams, _ := formatter.Format(cli.ApiRequest.Request.Params )
// 	// initialString := "https://1.1.1.1:666/api/diagnostics/firewall/log/?digest={digest}&limit={limit}"
// 	// digest := ""
// 	// formattedInitString, _ := formatter.Format(initialString, formatter.Named{
// 	// 	"digest": digest,
// 	// 	"limit":  "10000",
// 	// })
	

// }
