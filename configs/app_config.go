package configs

import (
	"errors"
	"os"
	"strconv"

	"github.com/tejashwinn/splitwise/types"

	_ "github.com/joho/godotenv/autoload"
)

func LoadConfig() (*types.Config, error) {
	jtAccessTokenExpMin, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXP_MIN"))
	if err != nil {
		return nil, errors.New("unable to parse jtAccessTokenExpMin")
	}
	jwtRefreshTokenExpHour, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXP_HOUR"))
	if err != nil {
		return nil, errors.New("unable to parse jwtRefreshTokenExpHour")
	}
	cfg := &types.Config{
		Server: types.ServerConfig{
			Port: os.Getenv("PORT"),
		},
		Db: types.DbConfig{
			Url: os.Getenv("DB_URL"),
		},
		Jwt: types.JwtConfig{
			SecretKey:              []byte(os.Getenv("JWT_SECRET_KEY")),
			Issuer:                 os.Getenv("JWT_ISSUER"),
			JwtAccessTokenExpMin:   int32(jtAccessTokenExpMin),
			JwtRefreshTokenExpHour: int32(jwtRefreshTokenExpHour),
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
