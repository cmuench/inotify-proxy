package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseValidYaml(t *testing.T) {

	validYamlData := `
---
watch:
- directory: /tmp/watch1
  extensions:
  -  ".scss"
  - ".js"
  - ".twig"

- directory: /tmp/watch2
  profile: magento2
  extensions:
  -  ".scss"
  - ".js"
  - ".twig"

`
	c, err := Parse([]byte(validYamlData))

	assert.NoError(t, err, "Config is valid and should not throw an error")
	assert.IsType(t, Config{}, c)

	assert.Equal(t, "/tmp/watch1", c.Entries[0].Directory)
	assert.Equal(t, "/tmp/watch2", c.Entries[1].Directory)

}

func TestParseInvalidYaml(t *testing.T) {
	invalidYamlData := `
---
watch

`
	_, err := Parse([]byte(invalidYamlData))

	assert.Error(t, err, "Config is invalid and should throw an error")
}

func TestLoad(t *testing.T) {

}
