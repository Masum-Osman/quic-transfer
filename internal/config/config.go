package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host        string
	Port        int
	FilePath    string
	Destination string
}

func Load(configType string) (*Config, error) {
	viper.SetConfigName(configType)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
