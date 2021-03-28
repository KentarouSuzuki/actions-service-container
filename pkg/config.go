package config

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Dialect		string `yaml:"dialect"`
    Datasource	string `yaml:"datasource"`
}

func NewConfig()(*DBConfig, error) {
	buf, err := ioutil.ReadFile("configs/dbconfig.yml")
	if err != nil {
		return nil, err
	}

	config := make(map[string]*DBConfig)
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return nil, err
	}

	environment, exists := os.LookupEnv("ENV")
	if !exists {
		environment = "development"
	}

	return config[environment], nil
}
