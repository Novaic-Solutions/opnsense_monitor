package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"github.com/Novaic-Solutions/opnsense_monitor/data"
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


//----------------------------------------------------------------------------
// Create response object based on the ApiRequest.ResponseObjType field.
//----------------------------------------------------------------------------
func (req *Request) CreateResponseObj(httpResp *http.Response) config.EndpointResponse {
	response := config.EndpointResponse{
		Uri: req.ApiRequest.Uri,
		Timestamp: time.Now().Format(time.RFC3339),
		ResponseDataType: req.ApiRequest.ResponseType,
		Data: nil,
	}

	byteArr, err := GetResponseData(httpResp)
	if err != nil {
		fmt.Printf("Request.go -- Error getting response data: %v\n", err)
		return response
	}

	var logEntry any

	switch req.ApiRequest.ResponseObjType {
	case "ArpTableEntry":
		logEntry = data.ArpTableEntry{}
	case "IFaceStatistics":
		logEntry = data.IfaceStatistics{}
	case "FirewallSession":
		logEntry = data.FirewallSession{}
	case "FirewallState":
		logEntry = data.FirewallState{}
	case "IfaceTraffic":
		logEntry = data.IfaceTraffic{}
	default:
		fmt.Printf("Request.go -- Unknown ResponseObjType: %s\n", req.ApiRequest.ResponseObjType)
	}

	if err := json.Unmarshal(byteArr, &logEntry); err != nil {
		fmt.Printf("Request.go -- Error unmarshaling response body to %s: %v\n", req.ApiRequest.ResponseObjType, err)
	}
	response.Data = logEntry
	
	return response
}



func (req *Request) Call() {

}