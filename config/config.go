package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config for the ms-login web service
var Config *MSCMConfig

func Initialize() {
	confContent, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	conf := &MSCMConfig{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		panic(err)
	}
}

// SetConfig bypasses the usual config override and validation behavior.
// It is intended to be used by unit tests.
func SetConfig(conf *MSCMConfig) {
	Config = conf
}

type MSCMConfig struct {
	Environment string `yaml:"environment"`
	Key         string `yaml:"key"`
}

const AppName = "ms-contact-manager"