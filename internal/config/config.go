package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Application struct {
		Name        string `yml:"name"`
		Version     string `yml:"version"`
		Author      string `yml:"author"`
		Description string `yml:"description"`
	} `yml:"application"`
}

func LoadConfig(filename string) (Config, error) {
	var cfg Config

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil

}

func GetVersion(cfg Config) string {
	return cfg.Application.Version
}
lu987456#