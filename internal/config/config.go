package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Watch struct {
	Dir string `yaml:"dir"`
}

type Config struct {
	Watch []Watch `yaml:"watch"`
	Profile string `yaml:"profile"`
}

func ReadFile(filename string) (Config, error) {
	var (
		c Config
		err error
		yamlData []byte
	)
	yamlData, err = ioutil.ReadFile(filename)

	if err != nil {
		return c, err
	}

	c, err = Parse(yamlData)

	return c, err
}

func Parse(yamlData []byte) (Config, error) {
	var c Config
	err := yaml.Unmarshal(yamlData, &c)

	if err != nil {
		return c, err
	}

	return c, nil
}
