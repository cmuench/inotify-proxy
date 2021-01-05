package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
)

// Config is the root struct for the application configuration
type Config struct {
	Entries []WatchEntry `yaml:"watch"`

	OldGlobalProfile *string `yaml:"profile"`
}

// WatchEntry is the configuration of one watch entry (directory) which is handled in a separate go-routine
type WatchEntry struct {
	Directory  string   `yaml:"directory"`
	Extensions []string `yaml:"extensions"`
	Profile    *string  `yaml:"profile"`
}

// AddEntry allows to add a new directory watch
func (c *Config) AddEntry(e WatchEntry) {
	c.Entries = append(c.Entries, e)
}

// GetEntryByDirectory returns the watch configuration of a given directory
func (c *Config) GetEntryByDirectory(dir string) WatchEntry {
	for _, e := range c.Entries {
		if e.Directory == dir {
			return e
		}
	}

	return WatchEntry{}
}

// Read a configuration from a file or other resource
func Read(f io.Reader) (Config, error) {
	var (
		c        Config
		err      error
		yamlData []byte
	)
	yamlData, err = ioutil.ReadAll(f)

	if err != nil {
		return c, err
	}

	c, err = Parse(yamlData)

	return c, err
}

// Parse the config data and return a Config object
func Parse(yamlData []byte) (Config, error) {
	var c Config
	err := yaml.Unmarshal(yamlData, &c)

	if err != nil {
		return c, err
	}

	return c, nil
}
