package config

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Settings ...
type Settings struct {
	Interface string `yaml:"interface"`
}

// Read reads a config file. Default file is config.yml
// A env CONFIG_PATH can be used instead.
// Fatal if can't parse config file
func Read() Settings {
	var path string
	if path = os.Getenv("CONFIG_PATH"); path == "" {
		path = "./config.yml"
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading config: %s", err.Error())
	}

	settings := Settings{}
	yaml.Unmarshal(data, &settings)

	if settings.Interface == "" {
		log.Fatal("Must set a valid interface on config file")
	}

	return settings
}
