package types

import (
	"time"
)

type Config struct {
	Db     DbConfig
	Server ServerConfig
	Jwt    JwtConfig
}

type JwtConfig struct {
	SecretKey              []byte
	Issuer                 string
	JwtAccessTokenExpMin   int32
	JwtRefreshTokenExpHour int32
}

type ServerConfig struct {
	Port string
}

type DbConfig struct {
	Url string
}

type User struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type TokenResposne struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Transaction struct {
	Id int64
}

type TransactionSplit struct {
	Id     int64
	UserId int64
	Share  float64
}
