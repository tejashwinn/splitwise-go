package types

type Config struct {
	Db     DbConfig
	Server ServerConfig
}

type ServerConfig struct {
	Port string
}

type DbConfig struct {
	Url string
}

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Transaction struct {
	Id int64
}

type TransactionSplit struct {
	Id     int64
	UserId int64
	Share  float64
}
