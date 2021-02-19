package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

// Config is a structure that represents trellostat config
type Config struct {
	Key   string
	Token string
	Board string
}

// ConfigFile parses the config file and returns config structure
func ConfigFile(filepath string) (Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, errors.New("error opening config file")
	}
	decoder := json.NewDecoder(bufio.NewReader(file))
	var c Config
	err = decoder.Decode(&c)
	if err != nil {
		return Config{}, errors.New("failing to decode the config file")
	}
	return c, nil
}
