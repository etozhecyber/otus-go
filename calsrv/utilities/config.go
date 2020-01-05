package utilities

import (
	"encoding/json"
	"os"
)

//Config model
type Config struct {
	HTTPListen string `json:"http_listen"`
	LogFile    string `json:"log_file"`
	LogLevel   string `json:"log_level"`
}

//GetConfiguration read from config file
func GetConfiguration(configPath string) (Config, error) {
	config := Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
