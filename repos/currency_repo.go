package repos

import (
	"context"

	"github.com/tejashwinn/splitwise/types"
)

type CurrencyRepo interface {
	FindAllOrderByName(ctx context.Context) ([]types.Currency, error)

	FindById(ctx context.Context, id int64) (*types.Currency, error)
}
