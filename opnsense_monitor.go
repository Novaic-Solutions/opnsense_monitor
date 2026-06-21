package main

import (
	"embed"
	"fmt"
	"os"
	"encoding/json"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"github.com/Novaic-Solutions/opnsense_monitor/client"
)

type ReqString struct {
        current int
        rowCount int
        sort map[string]string
        resolve string
}

//go:embed resources/config.yaml
var yamlFile embed.FS

func main() {
	fmt.Println("Starting application...")
	config := config.LoadConfig(yamlFile)
	fmt.Printf("Loaded config: %+v\n", config)

	api_key := config.Endpoint.ApiKey
	api_secret := config.Endpoint.ApiSecret

	req_string := ReqString{current: 1, rowCount: 50, sort: nil, resolve: "no"}
	jsonData, err := json.Marshal(req_string)
	if err != nil {
			fmt.Println("Json Serialization Error: ", err)
			os.Exit(1)
	}

	resp, err := client.SendRequest("POST", "https://192.168.1.1:666/api/diagnostics/interface/search_arp/", api_key, api_secret, jsonData)

	if err != nil {
			fmt.Printf("Error - performing request - info: %v\n", err)
			os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("Status Code: %v", resp.Status)


}


