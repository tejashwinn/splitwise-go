package types

import (
	"database/sql"
	"time"
)

type Transaction struct {
	Id int64
}

type TransactionSplit struct {
	Id     int64
	UserId int64
	Share  float64
}

type TransactionCurrency struct {
	ID           int          `json:"id"`
	Code         string       `json:"code"`
	Name         string       `json:"name"`
	Symbol       string       `json:"symbol"`
	ExchangeRate float64      `json:"exchangeRate"`
	BaseCurrency bool         `json:"baseCurrency"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    sql.NullTime `json:"updatedAt"`
}
