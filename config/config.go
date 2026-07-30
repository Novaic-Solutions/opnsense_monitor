package config

import (
	"embed"
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
	"github.com/Novaic-Solutions/opnsense_monitor/client"
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
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	API struct {
		ApiKey    string `yaml:"api_key"`
		ApiSecret string `yaml:"api_secret"`
		BaseURL   string `yaml:"base_url"`
		Endpoints []Endpoint `yaml:"endpoints"`
	} `yaml:"api"`
}




//----------------------------------------------------------------------
//       LoadConfig - Load the config.yaml file into a Config struct
//----------------------------------------------------------------------
func LoadConfig(yamlFile embed.FS) (*Config) {

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
