package config

import (
	"embed"
	"fmt"
	"os"
	//"reflect"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host	 string `yaml:"host"`
		Port	 int    `yaml:"port"`
		User	 string `yaml:"user"`
		Password string `yaml:"password"`
		Name	 string `yaml:"name"`
	} `yaml:"database"`
}


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
