package server

import (
	"fmt"
	"net/http"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"github.com/Novaic-Solutions/opnsense_monitor/client"
)

type Server struct {
	Port string
	Host string
	Conf *config.Config
	Clients []*client.Client
	ResponseChannel chan client.EndpointResponse
}

//----------------------------------------------------------------------------
//  Gather data from the database and form JSON response for the page
//----------------------------------------------------------------------------
func (serv Server) GatherData() string {
	totalClient := 0
	dataString := ""

	for _, cli := range serv.Clients {
		// Here, eventually, gather the data from the database for each client
		// and form the JSON response for the page.
		fmt.Printf("Gathering data from client: %v\n", cli.ApiRequest.Endpoint)
		go cli.Gather()
		totalClient++
	}

	for totalClient > 0 {
		response := <-serv.ResponseChannel
		dataString += response.Data + "\n"
		totalClient--
	}

	return dataString

}

//----------------------------------------------------------------------------
//  Request handler for all incoming requests.
//  This will serve up the JSON gathered from all of the API 
//  endpoints that are being monitored.
//----------------------------------------------------------------------------
func (serv Server) HandleAllRequest(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
	data := serv.GatherData()
    w.Write([]byte(data))
}

//----------------------------------------------------------------------------
//  Start the web server to serve the JSON data to the web page.
//----------------------------------------------------------------------------
func (serv Server) StartServer() {
	// Register the request handler for all incoming requests.
	http.HandleFunc("/", serv.HandleAllRequest)
	
	// Create the address string for the server to listen on.
	addr := fmt.Sprintf("%s:%s", serv.Host, serv.Port)
	
	// Start the server and listen for incoming requests.
	fmt.Printf("Starting server at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
