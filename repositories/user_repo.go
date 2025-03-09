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

	FindByEmailOrUsername(
		ctx context.Context,
		usernameEmail string,
	) (*types.User, error)
}
