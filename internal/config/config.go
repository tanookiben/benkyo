package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config ...
type Config struct {
	Addr string `json:"addr"`
}

var (
	defaultConfig = Config{
		Addr: ":8000",
	}
)

// Read ...
func Read() Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		return defaultConfig
	}
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	c := new(Config)
	if err := json.Unmarshal(raw, c); err != nil {
		log.Fatalf("Error unmarshaling config file: %v", err)
	}
	return *c
}
