package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type APIConfig struct {
	Apis map[string]APIInfo `yaml:"apis"`
}

type APIInfo struct {
	Method string `yaml:"method"`
	URL    string `yaml:"url"`
}

func LoadConfig(file string) (*APIConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := &APIConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return config, nil
}
