package main

import (
	"embed"
	"fmt"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"github.com/Novaic-Solutions/opnsense_monitor/client"
	// "github.com/Novaic-Solutions/opnsense_monitor/server"
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

	
	responseChannel := make(chan config.EndpointResponse, 100)
	fmt.Println("Starting application...")
	
	//-------------------------------------------------------------------------------
	// Load the configuration from the embedded config.yaml file.
	//-------------------------------------------------------------------------------
	conf := config.LoadConfig(yamlFile)
	fmt.Printf("Loaded config: %+v\n", conf)

	//-------------------------------------------------------------------------------
	// Create a slice to hold the clients. One for each endpoint in the config file.
	//-------------------------------------------------------------------------------
	requestClients, _ := client.PopulateApiRequests(conf.API.Endpoints, responseChannel, conf.API.ApiKey, conf.API.ApiSecret)
	fmt.Printf("Populated API requests: %+v\n", requestClients)

	//-------------------------------------------------------------------------------
	// Populate the requestClient slice with a Caller for each of the endpoints
	// in the config file.
	//-------------------------------------------------------------------------------
	
	// Send the client slice and the response channel to CreateApiRequests
	// To populate the slice with the clients for each of the endpoints in the config file.
	// err := CreateApiRequests(conf, &httpClients, responseChannel)
	// if err != nil {
	// 	fmt.Printf("Opnsense_monitor.go -- Error creating API requests: %v\n", err)
	// 	return
	// }

	//-------------------------------------------------------------------
	//	   Start client jobs for retrieving json from api endpoints
	//-------------------------------------------------------------------
	// Here, eventually, loop over the clients and start go routines for
	// each one.

	//-------------------------------------------------------------------
	//     Create the web server to serve the json data to the web page
	//-------------------------------------------------------------------
	// server := &server.Server{
	// 	Port:            conf.Server.Port,
	// 	Host:            conf.Server.Host,
	// 	Conf:            conf,
	// 	Clients:         httpClients,
	// 	ResponseChannel: responseChannel,
	// }

	//-------------------------------------------------------------------
	//     Start web server client to serve the json data to the web page
	//-------------------------------------------------------------------
	// server.StartServer()
}


