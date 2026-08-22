package main

import (
	"embed"
	"fmt"
	"sync"
	"os/signal"
	"context"
	"os"
	"syscall"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"github.com/Novaic-Solutions/opnsense_monitor/client"
	"github.com/Novaic-Solutions/opnsense_monitor/data"
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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	dataMutex := new(sync.Mutex)
	var requestClients []client.Caller
	responseChannel := make(chan config.EndpointResponse, 100)
	requestChannel := make(chan string, 5)
	outgoingChannel := make(chan string, 5)

	dataHandler := data.NewDataHandler(dataMutex, requestChannel, responseChannel, outgoingChannel)

	fmt.Println("Opnsense_monitor: Starting application...")

	//-------------------------------------------------------------------------------
	// Load the configuration from the embedded config.yaml file.
	//-------------------------------------------------------------------------------
	conf := config.LoadConfig(yamlFile)
	fmt.Printf("Opnsense_monitor: Loaded config: %+v\n", conf)

	//-------------------------------------------------------------------------------
	//	Start Data Handler routines.
	//-------------------------------------------------------------------------------
	go dataHandler.HandleIncomingData()
	go dataHandler.HandleRequests()

	//-------------------------------------------------------------------------------
	// Create a slice to hold the clients. One for each endpoint in the config file.
	//-------------------------------------------------------------------------------
	requestObject, _ := conf.CreateApiRequests()
	fmt.Printf("Opnsense_monitor: Populated API requests: %+v\n", requestObject)

	//-------------------------------------------------------------------------------
	// Create a client for each endpoint in the config file and start the client
	//-------------------------------------------------------------------------------
	for _, req := range requestObject {
		newClient := client.NewCaller(req, responseChannel)
		requestClients = append(requestClients, newClient)
	}

	//-------------------------------------------------------------------------------
	// Start the clients to call the endpoints and gather the data.
	//-------------------------------------------------------------------------------
	for _, client := range requestClients {
		fmt.Printf("Starting client for: %+v\n", client)
		go client.Call()
	}

	//-------------------------------------------------------------------
	//     Start web server client to serve the json data to the web page
	//-------------------------------------------------------------------
	server := &server.Server{
		Port:			conf.Server.Port,
		Host:			conf.Server.Host,
		Conf:			conf,
		RequestChannel:		requestChannel,
		OutgoingChannel:	outgoingChannel,
	}
	server.StartServer()
}
