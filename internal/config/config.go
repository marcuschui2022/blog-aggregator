package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(userHomeDir, configFileName)
	return fullPath, nil
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}

	var jsonConfig Config
	err = json.Unmarshal(file, &jsonConfig)
	if err != nil {
		return Config{}, err
	}

	return jsonConfig, nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonString, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(fullPath, jsonString, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	return write(*c)
}

//func write2(cfg Config) error {
//	fullPath, err := getConfigFilePath()
//	if err != nil {
//		return err
//	}
//
//	file, err := os.Create(fullPath)
//	if err != nil {
//		return err
//	}
//	defer file.Close()
//
//	encoder := json.NewEncoder(file)
//	err = encoder.Encode(cfg)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
