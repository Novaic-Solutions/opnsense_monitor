package config

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ----------------------------------------------------------------------
//
//	Structs for the config.yaml file objects
//
// ----------------------------------------------------------------------
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	API struct {
		ApiKey    string     `yaml:"api_key"`
		ApiSecret string     `yaml:"api_secret"`
		BaseURL   string     `yaml:"base_url"`
		Endpoints []Endpoint `yaml:"endpoints"`
	} `yaml:"api"`
}

// ----------------------------------------------------------------------
// Endpoint represents a single API Endpoint to call and the
// data required to perform the request.
// ----------------------------------------------------------------------
type Endpoint struct {
	Request         RequestObj `yaml:"request"`
	Type            string     `yaml:"type"`
	ResponseObjType string     `yaml:"response_obj_type"`
}

type RequestObj struct {
	Uri          string            `yaml:"uri"`
	Method       string            `yaml:"method"`
	ResponseType string            `yaml:"response_type"`
	Params       map[string]string `yaml:"params"`
	RequestBody  any               `yaml:"request_body"`
}

type EndpointResponse struct {
	Uri              string
	Timestamp        string
	ResponseDataType string
	Data             any
}

type ApiRequest struct {
	Url             string
	Uri             string
	Method          string
	Params          map[string]string
	Body            []byte
	ResponseType    string
	Username        string
	Password        string
	TypeRequest     string
	ResponseObjType string
}

// ----------------------------------------------------------------------
//
//	LoadConfig - Load the config.yaml file into a Config struct
//
// ----------------------------------------------------------------------
func LoadConfig(yamlFile embed.FS) *Config {

	//----------------------
	// Read file
	//----------------------
	fmt.Printf("Reading config file...\n")

	file, err := yamlFile.ReadFile("resources/config.yaml")
	if err != nil {
		fmt.Printf("Config.go -- Error - reading - config file: %v\n", err)
		os.Exit(1)
	}

	//----------------------
	// Create Config struct
	//----------------------
	config := Config{}

	//--------------------------------------------------------------------------------------------------
	// Unmarshal, which is their stupidass term for SERIALIZE or PARSE, the yaml file into the struct
	//--------------------------------------------------------------------------------------------------
	if err := yaml.Unmarshal(file, &config); err != nil {
		fmt.Printf("Config.go -- Error - parsing - config file: %v\n", err)
		os.Exit(1)
	}

	return &config
}

// ---------------------------------------------------------------------------
//
//	Populate the slice of clients with a client
//	for each endpoint in the config file.
//
// ---------------------------------------------------------------------------
func (conf *Config) CreateApiRequests() ([]*ApiRequest, error) {
	var apiRequests []*ApiRequest

	for _, endpoint := range conf.API.Endpoints {

		endp := endpoint.Request.RequestBody.(map[string]interface{})

		bytes, err := json.Marshal(endp)
		if err != nil {
			fmt.Printf("Config.go -- Error marshaling to JSON: %v", err)
		}

		// Prevents the body being sent as an empty JSON object when it is not needed for the request.
		// This should help prevent 400 errors.
		if len(endp) == 0 {
			bytes = nil
		}

		apiRequest := &ApiRequest{
			Url:             conf.API.BaseURL,
			Uri:             endpoint.Request.Uri,
			Method:          endpoint.Request.Method,
			Params:          endpoint.Request.Params,
			Body:            bytes,
			ResponseType:    endpoint.Request.ResponseType,
			Username:        conf.API.ApiKey,
			Password:        conf.API.ApiSecret,
			TypeRequest:     endpoint.Type,
			ResponseObjType: endpoint.ResponseObjType,
		}

		apiRequests = append(apiRequests, apiRequest)
	}

	return apiRequests, nil
}
