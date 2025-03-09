package repos

import (
	"context"

	"github.com/tejashwinn/splitwise/types"
)

type UserRepo interface {
	FindAll(
		ctx context.Context,
	) ([]types.User, error)

	Save(
		ctx context.Context,
		user *types.User,
	) (*types.User, error)

	FindByEmailOrUsername(
		ctx context.Context,
		usernameEmail string,
	) (*types.User, error)

	FindById(
		ctx context.Context,
		id int64,
	) (*types.User, error)
}
