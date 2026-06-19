package main

import (
	"embed"
	"fmt"
	"encoding/base64"
	"github.com/Novaic-Solutions/opnsense_monitor/config"
)

//go:embed resources/config.yaml
var yamlFile embed.FS

func basicAuth(username, password string) string {
  auth := username + ":" + password
  return base64.StdEncoding.EncodeToString([]byte(auth))
}


func main() {

	

	fmt.Println("Starting application...")
	config := config.LoadConfig(yamlFile)
	fmt.Printf("Loaded config: %+v\n", config)

}


