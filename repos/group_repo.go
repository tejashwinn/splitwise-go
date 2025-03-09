package repositories

import (
	"context"

	"github.com/tejashwinn/splitwise/types"
)

type GroupRepo interface {
	Save(ctx context.Context, user *types.Group) (*types.Group, error)
}
