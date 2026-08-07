package server

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
	"github.com/Novaic-Solutions/opnsense_monitor/client"
)

type Server struct {
	Port string
	Host string
	Conf *config.Config
	ResponseChannel chan *config.EndpointResponse
}

// //----------------------------------------------------------------------------
// //  Gather data from the database and form JSON response for the page
// //----------------------------------------------------------------------------
// func (serv Server) GatherData() string {
// 	totalClient := 0
// 	dataString := make(map[string]interface{})

// 	// Loop over the clients and start their loops in goroutines, where they
// 	// will 
// 	for _, cli := range serv.Clients {
// 		// Here, eventually, gather the data from the database for each client
// 		// and form the JSON response for the page.
// 		fmt.Printf("Server.go -- Gathering data from client: %v\n", cli.ApiRequest.Endpoint)
// 		go cli.Gather()
// 		totalClient++
// 	}

// 	// 
// 	for totalClient > 0 {
// 		response := <-serv.ResponseChannel
// 		if response.Data == nil {
// 			fmt.Printf("Server.go -- Error: No data received from client for URI: %s\n", response.Uri)
// 			totalClient--
// 			continue
// 		}
// 		dataString[response.Uri] = response.Data
// 		totalClient--
// 	}

// 	responseBytes, err := json.Marshal(dataString)
// 	if err != nil {
// 		fmt.Printf("Server.go -- Error marshaling data to JSON: %v\n", err)
// 		return ""
// 	}

// 	respString := string(responseBytes)

// 	return respString
// }

//----------------------------------------------------------------------------
//  Request handler for all incoming requests.
//  This will serve up the JSON gathered from all of the API 
//  endpoints that are being monitored.
//----------------------------------------------------------------------------
func (serv *Server) HandleAllRequest(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
	
	// Loop over the Data object that is updated by the server with
	// data from the clients gathered from the channel

    w.Write([]byte(data))
}

//----------------------------------------------------------------------------
//   Function for looping over the response channel and updating the Data object with
//   the data from the clients. This will be called in a goroutine.
//----------------------------------------------------------------------------
func (serv *Server) UpdateData(data *map[string]interface{}) {

}
//----------------------------------------------------------------------------
//  Start the web server to serve the JSON data to the web page.
//----------------------------------------------------------------------------
func (serv *Server) StartServer() {
	//------------------------------------------------------------------------
	// Create a pointer to a Data map that will hold the current data
	// gathered from the clients. This will be updated by the server with
	// data from the clients gathered from the channel.
	//------------------------------------------------------------------------
	var data map[string]interface{} = make(map[string]interface{})


	//------------------------------------------------------------------------
	// Register the request handler for all incoming requests.
	//------------------------------------------------------------------------
	http.HandleFunc("/", serv.HandleAllRequest)
	
	//------------------------------------------------------------------------
	// Create the address string for the server to listen on.
	//------------------------------------------------------------------------
	addr := fmt.Sprintf("%s:%s", serv.Host, serv.Port)
	
	//------------------------------------------------------------------------
	// Start the server and listen for incoming requests.
	//------------------------------------------------------------------------
	fmt.Printf("Server.go -- Starting server at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Server.go -- Error starting server: %v\n", err)
	}
}
