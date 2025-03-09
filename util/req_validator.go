package util

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
