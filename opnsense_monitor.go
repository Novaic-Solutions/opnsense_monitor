package main

import (
	"embed"
	"fmt"
	"os"
	//"encoding/json"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	//"github.com/Novaic-Solutions/opnsense_monitor/endpoint"
	"github.com/Novaic-Solutions/opnsense_monitor/client"
)

//go:embed resources/config.yaml
var yamlFile embed.FS

func main() {
	httpClients := make([]*client.Client, 0, 100)

	fmt.Println("Starting application...")
	conf := config.LoadConfig(yamlFile)
	fmt.Printf("Loaded config: %+v\n", conf)

	requests, err := client.PopulateApiRequests(conf)
	if err != nil {
		fmt.Printf("Error populating API requests: %v\n", err)
		os.Exit(1)
	}

	for _, req := range *requests {
		httpClients = append(httpClients, &client.Client{ApiRequest: &req})
	}

	// for _, cli := range httpClients {
	// 	cli.StartMonitoring()
	// }

	httpClients[0].StartMonitoring()

	// jsonData, err := json.Marshal(req_string)
	// if err != nil {
	// 		fmt.Println("Json Serialization Error: ", err)
	// 		os.Exit(1)
	// }

	// if err != nil {
	// 		fmt.Printf("Error - performing request - info: %v\n", err)
	// 		os.Exit(1)
	// }
	// defer resp.Body.Close()

	// fmt.Printf("Status Code: %v", resp.Status)

	//-------------------------------------------------------------------
	//	   Start client jobs for retrieving json from api endpoints
	//-------------------------------------------------------------------

	//-------------------------------------------------------------------
	//     
	//-------------------------------------------------------------------

	//-------------------------------------------------------------------
	//     Start web server client to serve the json data to the web page
	//-------------------------------------------------------------------

}


