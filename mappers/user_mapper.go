package mappers

import (
	"database/sql"
	"errors"

	"github.com/tejashwinn/splitwise/types"
)

func CreateReqToModel(req *types.UserReq) (*types.User, error) {
	return &types.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

func MapRowsToUser(rows *sql.Rows) (*types.User, error) {
	user := &types.User{}
	if err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return user, nil
}

func MapUserToUserRes(user *types.User) (*types.UserRes, error) {
	return &types.UserRes{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func MapUsersToUserRes(users []types.User) ([]types.UserRes, error) {
	usersRes := []types.UserRes{}
	for _, user := range users {
		userRes, err := MapUserToUserRes(&user)
		if err != nil {
			return nil, errors.New("Unable to map user")
		}
		usersRes = append(
			usersRes,
			*userRes,
		)
	}
	return usersRes, nil
}

func MapRowToUser(row *sql.Row) (*types.User, error) {
	user := &types.User{}
	if err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return user, nil
}
