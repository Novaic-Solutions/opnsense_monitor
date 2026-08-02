package client

 import (
 	"fmt"
	"net/http"
// 	"time"
// 	"io"
 	"encoding/json"
 	"github.com/Novaic-Solutions/opnsense_monitor/config"
// 	"gitlab.com/tymonx/go-formatter/formatter"
)

type FirewallLogStreamClient struct {
	ApiRequest *config.ApiRequest
	ResponseChannel chan config.EndpointResponse
}

func NewFirewallLogStreamClient(apiRequest *config.ApiRequest, responseChannel chan config.EndpointResponse) *FirewallLogStreamClient {
	return &FirewallLogStreamClient{
		ApiRequest: apiRequest,
		ResponseChannel: responseChannel,
	}
}

func


func (apiReq *FirewallLogStreamClient) Call() {
	var req *http.Request
	var err error

	url := apiReq.ApiRequest.Uri

	if len(apiReq.ApiRequest.Params) > 0 {
		formattedParams, _ := formatter.Format(apiReq.ApiRequest.Params)
		// 	formattedParams, _ := formatter.Format(cli.ApiRequest.Request.Params )
		// initialString := "https://1.1.1.1:666/api/diagnostics/firewall/log/?digest={digest}&limit={limit}"
		// digest := ""
		// formattedInitString, _ := formatter.Format(initialString, formatter.Named{
		// 	"digest": digest,
		// 	"limit":  "10000",
		// })
	
	}

	//---------------------------------------------------------------------------
	// Create a new request object with the appropriate method, endpoint, and body.
	// If the body is empty, set it to nil to prevent sending an empty JSON object.
	// This should help prevent 400 errors from the API.
	//---------------------------------------------------------------------------
	if len(apiReq.Body) == 0 {
		req, err = http.NewRequest(apiReq.Method, apiReq.Endpoint, nil)
	} else {
		//---------------------------------------------------------------------
		// Ensure only the content type for application/json is set IF
		// the request body is not empty. This prevents 400 errors from the API.
		//---------------------------------------------------------------------
		req, err = http.NewRequest(apiReq.Method, apiReq.Endpoint, bytes.NewBuffer(apiReq.Body))
		req.Header.Set("Content-Type", "application/json")
	}

	//---------------------------------------------------------------------------
	// Set the basic auth for the request using
	// the token created in opnsense.
	//---------------------------------------------------------------------------
	req.SetBasicAuth(apiReq.Username, apiReq.Password)
	
	//---------------------------------------------------------------------------
	// Set the basic auth for the request using
	// the token created in opnsense.
	//---------------------------------------------------------------------------
	req.SetBasicAuth(apiReq.Username, apiReq.Password)

	if err != nil {
		return nil, err
	}
}
