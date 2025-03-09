package config

import (
	"errors"
	"os"

	"github.com/tejashwinn/splitwise/types"

	_ "github.com/joho/godotenv/autoload"
)

func LoadConfig() (*types.Config, error) {
	cfg := &types.Config{
		Server: types.ServerConfig{
			Port: os.Getenv("PORT"),
		},
		Db: types.DbConfig{
			Url: os.Getenv("DB_URL"),
		},
	}

	if cfg.Server.Port == "" {
		return nil, errors.New("unable to find database url")
	}

	if cfg.Db.Url == "" {
		return nil, errors.New("unable to find database url")
	}

	return cfg, nil
}
