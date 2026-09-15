package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

// return the full path of the config-json
func getConfigFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// fulldir := homedir + "/" + configFileName
	fulldir := filepath.Join(homedir, configFileName)
	return fulldir, nil
}

// read ~/.gatorconfig.json, unmarshal return as Config struct
func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// changed os.readfile to os.open
	// data, err := os.ReadFile(filepath)
	data, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer data.Close()

	// rewritten to use decoder instead of json.unmarshal
	decoder := json.NewDecoder(data)
	conf := Config{}
	err = decoder.Decode(&conf)
	// err = json.Unmarshal(data, &conf)
	if err != nil {
		return Config{}, nil
	}

	return conf, nil
}

// write config file
func write(conf Config) error {
	filename, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// changed to os.create + json.encode from json.marhsal + os.writefile
	data, err := os.Create(filename)
	// data, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	defer data.Close()
	encoder := json.NewEncoder(data)
	err = encoder.Encode(conf)
	if err != nil {
		return err
	}
	//err = os.WriteFile(filename, data, 0666)
	return nil
}

func DefaultConfig() (Config, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	filename := homedir + "/" + configFileName

	data, err := json.Marshal(defConf)
	if err != nil {
		return Config{}, err
	}

	err = os.WriteFile(filename, data, 0666)
	if err != nil {
		return Config{}, err
	}
	return defConf, nil
}
