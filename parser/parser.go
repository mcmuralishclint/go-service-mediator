package parser

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Info     interface{} `yaml:"info"`
	Services interface{} `yaml:"services"`
}

func Parser() *Config {
	ymlData, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	configs := Config{}
	yaml.Unmarshal(ymlData, &configs)
	fmt.Println(configs)
	return &configs
}
