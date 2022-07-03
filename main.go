package main

import (
	"github.com/mcmuralishclint/go-service-mediator/parser"
)

var Config *parser.Config

func main() {
	Config = parser.Parser()
}
