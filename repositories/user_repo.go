package repositories

import (
	"context"

	"github.com/tejashwinn/splitwise/types"
)

type UserRepo interface {
	GetAllUsers(
		ctx context.Context,
	) ([]types.User, error)

	InsertOneUser(
		ctx context.Context,
		user *types.User,
	) (*types.User, error)

	FindByEmail(
		ctx context.Context,
		email string,
	) (*types.User, error)
}
