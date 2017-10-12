package config

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Current contains the result of the last call to `ReadAndUpdate` method
var Current Settings

// Settings ...
type Settings struct {
	Interface string `yaml:"interface"`

	Admin struct {
		Port uint   `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"admin"`
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

// Update read settings and update the Current var with its content
func Update() {
	Current = Read()
}
