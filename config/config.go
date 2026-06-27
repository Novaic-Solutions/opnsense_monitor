package config

import (
	"embed"
	"fmt"
	"os"
	//"reflect"
	"gopkg.in/yaml.v3"
)

//----------------------------------------------------------------------
//       Structs for the config.yaml file objects
//----------------------------------------------------------------------

type Config struct {
	Database struct {
		Host	 string `yaml:"host"`
		Port	 int    `yaml:"port"`
		User	 string `yaml:"user"`
		Password string `yaml:"password"`
		Name	 string `yaml:"name"`
	} `yaml:"database"`
	API struct {
		ApiKey    string `yaml:"api_key"`
		ApiSecret string `yaml:"api_secret"`
		BaseURL   string `yaml:"base_url"`
		Endpoints []Endpoint `yaml:"endpoints"`
	} `yaml:"api"`
}

// Endpoint represents a single API Endpoint to call and the 
// data required to perform the request.
type Endpoint struct {
	Uri string `yaml:"uri"`
	Method string `yaml:"method"`
	ResponseType string `yaml:"response_type"`
	Params string `yaml:"params"`
	RequestBody any `yaml:"request_body"`
}


//----------------------------------------------------------------------
//       LoadConfig - Load the config.yaml file into a Config struct
//----------------------------------------------------------------------
func LoadConfig(yamlFile embed.FS) (*Config) {

	// Read file
	fmt.Printf("Reading config file...\n")

	file, err := yamlFile.ReadFile("resources/config.yaml")
	if err != nil {
		fmt.Printf("Error - reading - config file: %v\n", err)
		os.Exit(1)
	}

	// // print the data from the config.yaml file
	// fmt.Printf("The data from the config.yaml file: %s\n", string(file))
	// fmt.Printf("The type of the data from the config.yaml file: %s\n", reflect.TypeOf(file))

	// Create Database struct
	config := Config{}

	// Unmarshal, which is their stupidass term for SERIALIZE or PARSE, the yaml file into the struct
	if err := yaml.Unmarshal(file, &config); err != nil {
		fmt.Printf("Error - parsing - config file: %v\n", err)
		os.Exit(1)
	}

	return &config
}
