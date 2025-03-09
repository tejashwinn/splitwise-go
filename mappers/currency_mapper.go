package mappers

import (
	"database/sql"

	"github.com/tejashwinn/splitwise/types"
)

func MapRowsToCurrency(rows *sql.Rows) (*types.Currency, error) {
	currency := &types.Currency{}
	if err := rows.Scan(
		&currency.Id,
		&currency.Code,
		&currency.Name,
		&currency.Symbol,
		&currency.ExchangeRate,
		&currency.BaseCurrency,
		&currency.CreatedAt,
		&currency.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return currency, nil
}

func MapRowToCurrency(row *sql.Row) (*types.Currency, error) {
	currency := &types.Currency{}
	if err := row.Scan(
		&currency.Id,
		&currency.Code,
		&currency.Name,
		&currency.Symbol,
		&currency.ExchangeRate,
		&currency.BaseCurrency,
		&currency.CreatedAt,
		&currency.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return currency, nil
}

func CurrencyModelToCurrencyRes(currency *types.Currency) (*types.CurrencyRes, error) {
	return &types.CurrencyRes{
		Id:     currency.Id,
		Code:   currency.Code,
		Name:   currency.Name,
		Symbol: currency.Symbol,
	}, nil

}
