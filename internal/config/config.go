package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ServerPort int `yaml:"server_port"`
}

func Load(file string) (*Config, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
