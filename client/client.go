package client

import (
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"fmt"
	"encoding/json"
)

type Client struct {
	
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

