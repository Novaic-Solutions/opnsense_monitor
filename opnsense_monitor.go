package main

import (
	"embed"
	"fmt"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
)

//go:embed resources/config.yaml
var yamlFile embed.FS

func main() {
	fmt.Println("Starting application...")
	config := config.LoadConfig(yamlFile)
	fmt.Printf("Loaded config: %+v\n", config)
}

// func main() {
// 	//---------------------------------------------------------------------------------
// 	//   Create a client with a custom timeout (Avoid http.DefaultClient)
// 	//---------------------------------------------------------------------------------
// 	client := &http.Client{
// 		Timeout: 10 * time.Second,
// 	}

// 	//---------------------------------------------------------------------------------
// 	//   Build a new request using a context
// 	//---------------------------------------------------------------------------------
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com", nil)
// 	if err != nil {
// 		fmt.Printf("Error creating request: %v\n", err)
// 		return
// 	}

// 	//---------------------------------------------------------------------------------
// 	//   Add custom headers if necessary
// 	//---------------------------------------------------------------------------------
// 	req.Header.Set("Accept", "application/json")
// 	req.Header.Set("User-Agent", "Go-HTTP-Client")

// 	//---------------------------------------------------------------------------------
// 	//   Execute the request
// 	//---------------------------------------------------------------------------------
// 	resp, err := client.Do(req)

// 	//---------------------------------------------------------------------------------
// 	//   resp will be the response object if the request was successful, or it will
// 	//   be nil if there was an error.
// 	//   err will be non-nil if there was an error sending the request or receiving the
// 	//   response. If err is non-nil, resp will be nil.
// 	//---------------------------------------------------------------------------------
// 	if err != nil {
// 		fmt.Printf("Error sending request: %v\n", err)
// 		return
// 	}

// 	//---------------------------------------------------------------------------------
// 	//   Close the response body to prevent resource leaks
// 	//---------------------------------------------------------------------------------
// 	defer resp.Body.Close()

// 	//---------------------------------------------------------------------------------
// 	//   Check the response status code
// 	//---------------------------------------------------------------------------------
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
// 		return
// 	}

// 	//---------------------------------------------------------------------------------
// 	//   Read the response payload
// 	//---------------------------------------------------------------------------------
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Printf("Error reading response: %v\n", err)
// 		return
// 	}

// 	fmt.Println(string(body))
// }
