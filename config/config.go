package config

import (
	"embed"
	"fmt"
	//"os"
	//"gopkg.in/yaml.v3"
)

type Database struct {
		Host	 string `yaml:"host"`
		Port	 int    `yaml:"port"`
		User	 string `yaml:"user"`
		Password string `yaml:"password"`
		Name	 string `yaml:"name"`
}

func LoadConfig(yamlFile embed.FS) (*Database) {
	// Read file
	fmt.Printf("Reading config file...\n")
	// file, err := yamlFile.ReadFile("resources/config.yaml")
	// if err != nil {
	// 	fmt.Printf("Error 2222222 reading config file: %v\n", err)
	// 	os.Exit(1)
	// }

	// Create Database struct
	config := &Database{}

	// Unmarshal, which is their stupidass term for SERIALIZE or PARSE, the yaml file into the struct
	// if err := yaml.Unmarshal(file, config); err != nil {
	// 	fmt.Printf("Error 1111111 parsing config file: %v\n", err)
	// 	os.Exit(1)
	// }

	return config
}
