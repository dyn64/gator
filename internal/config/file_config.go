package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

// return the full path of the config-json
func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fulldir := homedir + "/" + configFileName
	return fulldir, nil
}

// write config file
func write(conf Config) error {
	filename, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0666)
	return nil
}

func defaultConfig() error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filename := homedir + "/" + configFileName
	defConf := Config{
		DbURL:           "",
		CurrentUserName: "",
	}
	configFile, err := os.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}
}
