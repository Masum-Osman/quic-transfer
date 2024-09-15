package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	FilePath    string `yaml:"filePath"`
	Destination string `yaml:"destination"`
}

func Load(configType string) (*Config, error) {
	configPath := filepath.Join("config", configType+".yaml")
	fmt.Printf("Attempting to read config file: %s\n", configPath)

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	fmt.Println("File contents:")
	fmt.Println(string(data))

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	fmt.Printf("Parsed config: %+v\n", cfg)

	return &cfg, nil
}
