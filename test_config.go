package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func main() {
	// Read the YAML file
	data, err := ioutil.ReadFile("internal/config/server.yaml")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the contents of the file
	fmt.Println("File contents:")
	fmt.Println(string(data))

	// Try to parse the YAML
	var result map[string]interface{}
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("Error parsing YAML: %v\n", err)
		return
	}

	fmt.Println("Parsed YAML:")
	fmt.Printf("%+v\n", result)
}
