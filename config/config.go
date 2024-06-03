package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	AWS struct {
		TableName string `toml:"TableName"`
	} `toml:"aws"`
}

func LoadConfig(filePath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
