package client

import (
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Client struct {
	ApiRequest *ApiRequest
	// Probably a channel to send response back on
}

func PopulateApiRequests(conf *config.Config) (*[]ApiRequest, error) {
	apiObj := make([]ApiRequest, 0, 100)

	for _, value := range conf.API.Endpoints {
		fmt.Println("---------------------------------")
		//fmt.Println(value)
		//fmt.Println(reflect.TypeOf(value.RequestBody))
		endp := value.RequestBody.(map[string]interface{})

		bytes, err := json.Marshal(endp)
		if err != nil {
			fmt.Printf("Error marshaling to JSON: %v", err)
		}

		newReq := ApiRequest{
			Method: value.Method,
			Endpoint: conf.API.BaseURL + value.Uri,
			ResponseType: value.ResponseType,
			Username: conf.API.ApiKey,
			Password: conf.API.ApiSecret,
			Body: bytes,
		}
		apiObj = append(apiObj, newReq)
	}

	if len(apiObj) == 0 {
		return nil, fmt.Errorf("No API requests found in configuration")
	}

	return &apiObj, nil
}

func (cli *Client) StartMonitoring() {
	resp, err := cli.ApiRequest.SendRequest()
	if err != nil {
		fmt.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v", err)
		return
	}

	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", string(body))

	// Add logic here to process the response and send it through the channel
	// for further processing

}

