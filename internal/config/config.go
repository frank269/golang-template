package config

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

func init() {
	// load environment variables from .env file if it exists
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}
}

func LoadConfig() (*AppConfig, error) {
	config := &AppConfig{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	dbconfig := &DatabaseConfig{}
	if err := env.Parse(dbconfig); err != nil {
		return nil, err
	}
	config.DatabaseConfig = dbconfig
	return config, nil
}
