package mappers

import (
	"database/sql"

	"github.com/tejashwinn/splitwise/types"
)

func CreateReqToModel(req *types.UserReq) (*types.User, error) {
	return &types.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

func MapUserRows(rows *sql.Rows) (*types.User, error) {
	user := &types.User{}
	if err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	); err != nil {
		return nil, err
	}
	return user, nil
}

func MapUserToUserRe(user *types.User) (*types.UserRes, error) {
	return &types.UserRes{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func MapUserRow(row *sql.Row) (*types.User, error) {
	user := &types.User{}
	if err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	); err != nil {
		return nil, err
	}
	return user, nil
}
