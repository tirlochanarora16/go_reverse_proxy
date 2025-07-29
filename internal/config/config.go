package config

import (
	"fmt"
	"os"

	"github.com/tirlochanarora16/go_reverse_proxy/internal/lb"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Routes []Route `yaml:"routes"`
}

type Route struct {
	Path      string     `yaml:"path"`
	Target    string     `yaml:"target"`
	RateLimit *RateLimit `yaml:"rate_limit,omitempty"`
}

type RateLimit struct {
	Rate  int `yaml:"rate"`
	Burst int `yaml:"burst"`
}

var ConfigFileData Config

func ParseConfigFile() {
	fileResult, err := os.ReadFile(lb.ConfigFileName)

	if err != nil {
		fmt.Printf("Erro reading the config file %s", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(fileResult, &ConfigFileData)

	if err != nil {
		fmt.Println("Error unmarshing the json file")
		os.Exit(1)
	}
}
