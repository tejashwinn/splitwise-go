package types

import (
	"database/sql"
	"time"
)

type Group struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	CurrencyId  int64         `json:"currencyId"`
	CreatedAt   time.Time     `json:"createdAt"`
	CreatedBy   int64         `json:"createdBy"`
	UpdatedAt   sql.NullTime  `json:"updatedAt"`
	UpdatedBy   sql.NullInt64 `json:"updatedBy"`
}

type GroupReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CurrencyId  int64  `json:"currencyId"`
}

type GroupRes struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Currency    CurrencyRes `json:"currencyId"`
	CreatedBy   UserRes     `json:"createdBy"`
	CreatedAt   time.Time   `json:"createdAt"`
}
