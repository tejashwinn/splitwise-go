package repositories

import (
	"context"

	"github.com/tejashwinn/splitwise/types"
)

// UserRepository defines methods for user data operations
type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]types.User, error)
}
