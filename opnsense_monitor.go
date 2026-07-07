package main

import (
	"embed"
	"fmt"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"github.com/Novaic-Solutions/opnsense_monitor/client"
	"github.com/Novaic-Solutions/opnsense_monitor/server"
)

//----------------------------------------------------------------------------
// Initially going to leave the clients and server in the main package,
// because eventually I may add functionality to have the clients
// continually monitoring the endpoints and storing the data in a database,
// and then have the server serve that data to the web page.
//
// As it stands now, the client will gather the data from the endpoints once
// the page is accessed that is being served up from the server.
//----------------------------------------------------------------------------

//go:embed resources/config.yaml
var yamlFile embed.FS


//------------------------------------------------------------------------------
//	 Create the request objects for each of the API endpoints
//------------------------------------------------------------------------------
func CreateApiRequests(conf *config.Config, httpClients *[]*client.Client, responseChannel chan client.EndpointResponse) error {
	//-------------------------------------------------------------------
	//     Create the request objects for each of the API endpoints
	//     then create a client for each of the requests and start 
	// 	   monitoring
	//-------------------------------------------------------------------
	requests, err := client.PopulateApiRequests(conf)
	if err != nil {
		fmt.Printf("Error populating API requests: %v\n", err)
		return err
	}

	//-------------------------------------------------------------------
	//    Create a client for each of the requests and start monitoring
	//-------------------------------------------------------------------
	for _, req := range *requests {
		*httpClients = append(*httpClients, &client.Client{ApiRequest: &req, ResponseChannel: responseChannel})
	}
	return nil
}

//----------------------------------------------------------------------------
//	Initialize the application, load the configuration, create the API requests,
//  and start the monitoring clients.
//----------------------------------------------------------------------------
func init() {
	// Good for setting up database connections, etc.
	// Not used yet in this application, but could be used in the future.
}

//----------------------------------------------------------------------------
//	Main entry point for the application.
//----------------------------------------------------------------------------
func main() {
	httpClients := make([]*client.Client, 0, 100)
	responseChannel := make(chan client.EndpointResponse, 100)

	fmt.Println("Starting application...")
	conf := config.LoadConfig(yamlFile)
	fmt.Printf("Loaded config: %+v\n", conf)

	err := CreateApiRequests(conf, &httpClients, responseChannel)
	if err != nil {
		fmt.Printf("Error creating API requests: %v\n", err)
		return
	}

	//-------------------------------------------------------------------
	//	   Start client jobs for retrieving json from api endpoints
	//-------------------------------------------------------------------
	// Here, eventually, loop over the clients and start go routines for
	// each one.

	//-------------------------------------------------------------------
	//     Create the web server to serve the json data to the web page
	//-------------------------------------------------------------------
	server := &server.Server{
		Port:            conf.Server.Port,
		Host:            conf.Server.Host,
		Conf:            conf,
		Clients:         httpClients,
		ResponseChannel: responseChannel,
	}

	//-------------------------------------------------------------------
	//     Start web server client to serve the json data to the web page
	//-------------------------------------------------------------------
	server.StartServer()
}


