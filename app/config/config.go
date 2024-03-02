package config

import (
	"mlogreport/utils/validation"
)

type Config struct {
	DBHOST      string `env:"DBHOST"`
	DBPORT      int    `env:"DBPORT"`
	DBUSER      string `env:"DBUSER"`
	DBPASS      string `env:"DBPASS"`
	DBNAME      string `env:"DBNAME"`
	SERVERPORT  string `env:"SERVERPORT"`
	API_STORAGE string `env:"API_STORAGE"`
	MODE        string `env:"MODE"`
}

func InitConfig() *Config {
	config := &Config{}
	validation.EnvCheck(config)
	return config
}
