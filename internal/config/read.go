package config

import (
	"encoding/json"
	"os"
)

// read ~/.gatorconfig.json, unmarshal return as Config struct
func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}

	conf := Config{}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return Config{}, nil
	}

	return conf, nil
}
