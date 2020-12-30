package config

import (
	"gopkg.in/yaml.v3"
	"io"
)

type Config struct {
	Entries []WatchEntry `yaml:"watch"`
}

type WatchEntry struct {
	Directory  string   `yaml:"directory"`
	Extensions []string `yaml:"extensions"`
	Profile    *string  `yaml:"profile"`
}

func (c *Config) AddEntry(e WatchEntry) {
	c.Entries = append(c.Entries, e)
}

func (c *Config) GetEntryByDirectory(dir string) WatchEntry {
	for _, e := range c.Entries {
		if e.Directory == dir {
			return e
		}
	}

	return WatchEntry{}
}

func Read(f io.Reader) (Config, error) {
	var (
		c        Config
		err      error
		yamlData []byte
	)
	_, err = f.Read(yamlData)

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
