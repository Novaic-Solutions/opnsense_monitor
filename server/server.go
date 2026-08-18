package server

import (
	"fmt"
	"net/http"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
)

type Server struct {
	Port string
	Host string
	Conf *config.Config
	RequestChannel chan string
	OutgoingChannel chan string
}

//----------------------------------------------------------------------------
//  Request handler for all incoming requests.
//  This will serve up the JSON gathered from all of the API 
//  endpoints that are being monitored.
//----------------------------------------------------------------------------
func (serv *Server) HandleAllRequest(w http.ResponseWriter, r *http.Request) {
    responseBytes := []byte(serv.GetData())
	w.WriteHeader(http.StatusOK)
    w.Write(responseBytes)
}

func (serv *Server) GetData() string {
	serv.RequestChannel <- "metrics"
	return <-serv.OutgoingChannel
}

//----------------------------------------------------------------------------
//  Start the web server to serve the JSON data to the web page.
//----------------------------------------------------------------------------
func (serv *Server) StartServer() {

	//------------------------------------------------------------------------
	// Register the request handler for all incoming requests.
	//------------------------------------------------------------------------
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serv.HandleAllRequest(w, r)
	})
	
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
