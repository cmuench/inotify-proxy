package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseValidYaml(t *testing.T) {

	validYamlData := `
---
watch:
  - dir: /tmp/watch1
  - dir: /tmp/watch2
profile: magento2

`
	c, err := Parse([]byte(validYamlData))

	assert.NoError(t, err, "Config is valid and should not throw an error")
	assert.IsType(t, Config{}, c)
}

func TestParseInvalidYaml(t *testing.T) {
	invalidYamlData := `
---
watch

`
	_, err := Parse([]byte(invalidYamlData))

	assert.Error(t, err, "Config is invalid and should throw an error")
}
