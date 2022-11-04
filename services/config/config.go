package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}

var (
	Config *ConfigStruct
)

func InitConfig(configPath string) (*ConfigStruct, error) {
	config := &ConfigStruct{}

	file, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil

}

func SetPublicConfig(cfg *ConfigStruct) {
	Config = cfg
}
