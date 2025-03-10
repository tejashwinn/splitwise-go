package repos

import (
	"context"

	"github.com/tejashwinn/splitwise/types"
)

type GroupUserRepo interface {
	Save(ctx context.Context, user *types.GroupUser) (*types.GroupUser, error)

	FindUserIdByGroupId(ctx context.Context, groupId int64) ([]int64, error)
}
