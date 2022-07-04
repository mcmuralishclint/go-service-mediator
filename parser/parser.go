package parser

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Info     interface{} `yaml:"info"`
	Services []map[string]Service
}

type Service struct {
	Name    string                  `yaml:"name"`
	BaseUrl string                  `yaml:"base_url"`
	Version []map[string][]Endpoint `yaml:"version"`
}

type Endpoint struct {
	Name         string `yaml:"name"`
	Url          string `yaml:"url"`
	Verb         string `yaml:"verb"`
	ContentType  string `yaml:"content-type"`
	EndpointType string `yaml:"endpoint-type"`
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
