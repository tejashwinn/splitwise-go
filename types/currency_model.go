package types

import (
	"database/sql"
	"time"
)

type Currency struct {
	Id           int64        `json:"id"`
	Code         string       `json:"code"`
	Name         string       `json:"name"`
	Symbol       string       `json:"symbol"`
	ExchangeRate float64      `json:"exchangeRate"`
	BaseCurrency bool         `json:"baseCurrency"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    sql.NullTime `json:"updatedAt"`
}

type CurrencyRes struct {
	Id     int64  `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
