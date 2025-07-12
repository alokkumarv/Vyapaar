package config

import "os"

type Config struct {
	DB string
}

func GetConfig() *Config {
	return &Config{
		DB: os.Getenv("DATABASE"),
	}
}
