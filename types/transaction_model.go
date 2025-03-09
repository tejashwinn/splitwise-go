package types

type Transaction struct {
	Id int64
}

type TransactionSplit struct {
	Id     int64
	UserId int64
	Share  float64
}
