package urlzap

import (
	"io"

	"gopkg.in/yaml.v2"
)

// HTTP configs for HTTP server
type HTTP struct {
	BasePath string `yaml:"basePath" json:"basePath"`
	Address  string `yaml:"address" json:"address"`
}

// URLs define a key to value map with the URLs.
// - Keys are the value to be mapped by the service and the values are the URLs to be redirected to.
// - Nested keys would have a "/" for each nesting level.
type URLs map[interface{}]interface{}

// Config define config payload.
type Config struct {
	Path string `yaml:"path" json:"path" default:"./" envconfig:"PATH"`
	HTTP HTTP   `yaml:"http" json:"http"`
	URLs URLs   `yaml:"urls" json:"urls"`
}

// FromYAML loads config from a yaml.
func FromYAML(r io.Reader) (Config, error) {
	c := Config{}
	if err := yaml.NewDecoder(r).Decode(&c); err != nil {
		return c, err
	}

	return c, nil
}
