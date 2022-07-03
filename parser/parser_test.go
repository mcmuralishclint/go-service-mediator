package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var filename string = "./config_test.yaml"

func TestParser(t *testing.T) {
	config := Parser(filename)
	assert.NotNil(t, config, "Config should not be nil")
	assert.NotNil(t, &config.Info, "Config Info should not be nil")
	assert.NotNil(t, &config.Services, "Config Services should not be nil")
}
