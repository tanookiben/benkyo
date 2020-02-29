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

// Read ...
func Read() Config {
	raw, err := ioutil.ReadFile(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	c := new(Config)
	if err := json.Unmarshal(raw, c); err != nil {
		log.Fatalf("Error unmarshaling config file: %v", err)
	}
	return *c
}
