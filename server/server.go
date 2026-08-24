package server

import (
	"fmt"
	"net/http"
	"errors"
	"time"
	"os"
	"os/signal"
	"syscall"
	"context"
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
	fmt.Printf("Server.go -- Received request: %s\n", r.URL.Path)

    responseBytes := []byte(serv.GetData())
	w.WriteHeader(http.StatusOK)
    w.Write(responseBytes)
	fmt.Printf("Server.go -- Sent response: %d bytes\n", len(responseBytes))
}

func (serv *Server) GetData() string {
	serv.RequestChannel <- "metrics"
	return <-serv.OutgoingChannel
}

//----------------------------------------------------------------------------
//  Start the web server to serve the JSON data to the web page.
//----------------------------------------------------------------------------
func (serv *Server) StartServer(ctx context.Context) {
	httpServer := &http.Server{
		Addr: fmt.Sprintf("%s:%s", serv.Host, serv.Port),
	}

	//------------------------------------------------------------------------
	// Register the request handler for all incoming requests.
	//------------------------------------------------------------------------
	http.HandleFunc("/", serv.HandleAllRequest)
	
	//------------------------------------------------------------------------
	// Create the address string for the server to listen on.
	//------------------------------------------------------------------------
	
	//------------------------------------------------------------------------
	// Start the server and listen for incoming requests.
	//------------------------------------------------------------------------
	fmt.Printf("Server.go -- Starting server at %s\n", httpServer.Addr)
	go func() {
		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server.go -- server error: %v\n", err)
		}
	}()

	<-ctx.Done()

    shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
    defer shutdownRelease()

    if err := httpServer.Shutdown(shutdownCtx); err != nil {
        fmt.Printf("Server.go -- HTTP shutdown error: %v\n", err)
    }
    fmt.Printf("Server.go -- Graceful shutdown complete.\n")
}
