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

func ReadFile(filename string) Config {
	yamlData, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var c Config
	err = yaml.Unmarshal(yamlData, &c)

	if err != nil {
		panic(err)
	}

	return c
}

