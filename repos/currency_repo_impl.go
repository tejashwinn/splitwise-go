package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tejashwinn/splitwise/mappers"
	"github.com/tejashwinn/splitwise/types"
)

type CurrencyRepoImpl struct {
	DB *sql.DB
}

func NewCurrencyRepository(db *sql.DB) CurrencyRepo {
	return &CurrencyRepoImpl{DB: db}
}

func (repo *CurrencyRepoImpl) FindAllOrderByName(ctx context.Context) ([]types.Currency, error) {
	query := `
		SELECT OBJECT_ID,
			CUR_CODE,
			CUR_NAME,
			CUR_SYMBOL,
			CUR_EX_RATE,
			CUR_BASE_YN,
			CREATED_AT,
			UPDATED_AT
		FROM SW_CUR
		ORDER BY CUR_NAME
	`
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	currencies := []types.Currency{}
	for rows.Next() {
		currency, err := mappers.MapRowsToCurrency(rows)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, *currency)
	}
	return currencies, nil
}

func (repo *CurrencyRepoImpl) FindById(ctx context.Context, id int64) (*types.Currency, error) {
	query := `
		SELECT OBJECT_ID,
			CUR_CODE,
			CUR_NAME,
			CUR_SYMBOL,
			CUR_EX_RATE,
			CUR_BASE_YN,
			CREATED_AT,
			UPDATED_AT
		FROM SW_CUR
		WHERE OBJECT_ID = $1
	`
	row := repo.DB.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		return nil, errors.New("Unable to find currency")
	}
	currency, err := mappers.MapRowToCurrency(row)
	if err != nil {
		return nil, err
	}
	return currency, nil
}
