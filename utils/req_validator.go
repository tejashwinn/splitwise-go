package utils

import (
	"errors"

	"github.com/tejashwinn/splitwise/types"
)

func ValidateCreateUser(req *types.UserReq) error {
	if len(req.Name) > 255 {
		return errors.New("title length cannot be greater than 100 characters")
	}
	if len(req.Email) > 255 {
		return errors.New("content length cannot be greater than 1000 characters")
	}
	if len(req.Password) > 255 {
		return errors.New("content length cannot be greater than 1000 characters")
	}
	return nil
}

func ValidateLoginUser(req *types.LoginReq) error {
	if req.User == "" {
		return errors.New("email/username cannot be empty")
	}
	if req.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}
