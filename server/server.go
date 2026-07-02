package server

import (
	"fmt"
	"net/http"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
)

type Server struct {
	Port string
	Host string
	Conf *Config
}

//----------------------------------------------------------------------------
//  Gather data from the database and form JSON response for the page
//----------------------------------------------------------------------------
func (serv Server) GatherData() {
}

//----------------------------------------------------------------------------
//  Request handler for all incoming requests.
//  This will serve up the JSON gathered from all of the API 
//  endpoints that are being monitored.
//----------------------------------------------------------------------------
func (serv Server) HandleAllRequest(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
	// data := serv.GatherData()
    w.Write([]byte(fmt.Sprintf("Response from %v\n", service.Name)))
}

//----------------------------------------------------------------------------
//  Start the web server to serve the JSON data to the web page.
//----------------------------------------------------------------------------
func (serv Server) StartServer() {

}
