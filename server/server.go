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
	ResponseChannel chan config.EndpointResponse
}

//----------------------------------------------------------------------------
//  Request handler for all incoming requests.
//  This will serve up the JSON gathered from all of the API 
//  endpoints that are being monitored.
//----------------------------------------------------------------------------
func (serv *Server) HandleAllRequest(w http.ResponseWriter, r *http.Request, data *map[string]string) {
    w.WriteHeader(http.StatusOK)
	
	// Loop over the Data object that is updated by the server with
	// data from the clients gathered from the channel

    // w.Write([]byte(data))
}

//----------------------------------------------------------------------------
//   Function for looping over the response channel and updating the Data object with
//   the data from the clients. This will be called in a goroutine.
//----------------------------------------------------------------------------
func (serv *Server) UpdateData(data *map[string]string) {
	// for responseData := range serv.ResponseChannel {
	// 	// Update the Data object with the data from the clients gathered from the channel

	// }
}

//----------------------------------------------------------------------------
//   Function for formatting the data from the clients into a string
//   so that it conforms to the output that is expected by prometheus.
//----------------------------------------------------------------------------
func (serv *Server) FormatDataString(data *map[string]string) string {
	return ""
}

//----------------------------------------------------------------------------
//   Get the data from the response object
//----------------------------------------------------------------------------
func (serv *Server) GetDataFromResponse(response config.EndpointResponse) string {
	switch response.ResponseDataType {
	case "FirewallLogEntry":
		// Format the data from the response object into a string
		return ""
	case "ArpTableEntry":
		// Format the data from the response object into a string
		return ""
	case "IFaceStatistics":
		// Format the data from the response object into a string
		return ""
	case "FirewallSession":
		// Format the data from the response object into a string
		return ""
	case "FirewallState":
		// Format the data from the response object into a string
		return ""
	case "IfaceTraffic":
		// Format the data from the response object into a string
		return ""
		
	default:
		fmt.Printf("Server.go -- Unknown ResponseDataType: %s\n", response.ResponseDataType)
		return ""
	}
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
	var data map[string]string = make(map[string]string)
	var dataPtr *map[string]string = &data

	//------------------------------------------------------------------------
	// Register the request handler for all incoming requests.
	//------------------------------------------------------------------------
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serv.HandleAllRequest(w, r, dataPtr)
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
