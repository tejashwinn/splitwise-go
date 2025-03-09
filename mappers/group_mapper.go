package mappers

import (
	"github.com/tejashwinn/splitwise/types"
)

func CreateReqToGroupModel(req *types.GroupReq) (*types.Group, error) {
	return &types.Group{
		Name:        req.Name,
		Description: req.Description,
		CurrencyId:  req.CurrencyId,
	}, nil
}

func GroupModelToGroupRes(
	group *types.Group,
	user *types.User,
	currency *types.Currency,
) (*types.GroupRes, error) {
	userRes, err := MapUserToUserRes(user)
	if err != nil {
		return nil, err
	}
	currencyRes, err := CurrencyModelToCurrencyRes(currency)
	if err != nil {
		return nil, err
	}
	return &types.GroupRes{
		Id:          group.Id,
		Name:        group.Name,
		Description: group.Description,
		Currency:    *currencyRes,
		CreatedBy:   *userRes,
		CreatedAt:   group.CreatedAt,
	}, nil
}
