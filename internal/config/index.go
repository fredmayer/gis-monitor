package config

import (
	"flag"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"os"
)

var Instance *Config

func MustLoad() *Config {
	if Instance != nil {
		return Instance
	}
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	Instance = MustLoadPath(configPath)
	return Instance
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	config.AddDriver(yaml.Driver)
	err := config.LoadFiles(configPath)
	if err != nil {
		panic(err)
	}

	config.Decode(&cfg)

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
