package parser

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Info     interface{} `yaml:"info"`
	Services interface{} `yaml:"services"`
}

func Parser(filename string) *Config {
	ymlData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	configs := Config{}
	yaml.Unmarshal(ymlData, &configs)
	return &configs
}
