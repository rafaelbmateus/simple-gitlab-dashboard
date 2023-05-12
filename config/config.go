package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name   string `yaml:"name"`
	GitLab struct {
		API  string `yaml:"api"`
		Host string `yaml:"host"`
	}
	Projects []string `yaml:"projects"`
	Users    []string `yaml:"users"`
}

// Load configuration file.
func Load(configFile string) (*Config, error) {
	cfg, err := readConfigurationFile(configFile)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// readConfigurationFile read configuration file.
func readConfigurationFile(fileName string) (config *Config, err error) {
	var bytes []byte
	if bytes, err = os.ReadFile(fileName); err == nil {
		return parseAndValidateConfig(bytes)
	}
	return
}

// parseAndValidateConfig parse configuration file and validates.
func parseAndValidateConfig(yamlBytes []byte) (config *Config, err error) {
	yamlBytes = []byte(os.ExpandEnv(string(yamlBytes)))
	if err = yaml.Unmarshal(yamlBytes, &config); err != nil {
		return
	}

	return
}
