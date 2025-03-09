package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/tejashwinn/splitwise/constants"
	"github.com/tejashwinn/splitwise/mappers"
	"github.com/tejashwinn/splitwise/repos"
	"github.com/tejashwinn/splitwise/types"
	"github.com/tejashwinn/splitwise/utils"
)

type GroupHandler struct {
	UserRepo     repos.UserRepo
	GroupRepo    repos.GroupRepo
	CurrencyRepo repos.CurrencyRepo
	JwtUtil      utils.JwtUtil
}

func NewGroupHandler(
	userRepo repos.UserRepo,
	groupRepo repos.GroupRepo,
	currencyRepo repos.CurrencyRepo,
	jwtUtil *utils.JwtUtil,
) *GroupHandler {
	return &GroupHandler{UserRepo: userRepo, GroupRepo: groupRepo, CurrencyRepo: currencyRepo, JwtUtil: *jwtUtil}
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	req := &types.GroupReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}

	group, err := mappers.CreateReqToGroupModel(req)
	if err != nil {
		http.Error(w, "Unable to map request", http.StatusBadRequest)
		return
	}
	group, err = h.GroupRepo.Save(r.Context(), group)
	if err != nil {
		http.Error(w, "Unable to save request", http.StatusInternalServerError)
		return
	}
	groupRes, err := h.groupToGroupRes(group, w, r)
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(groupRes)
}

func (h *GroupHandler) groupToGroupRes(
	group *types.Group,
	w http.ResponseWriter, r *http.Request,
) (*types.GroupRes, error) {

	user, err := h.UserRepo.FindById(context.Background(), group.CreatedBy)
	if err != nil {
		return nil, errors.New("User not found")
	}
	currency, err := h.CurrencyRepo.FindById(context.Background(), group.CurrencyId)
	if err != nil {
		return nil, errors.New("User not found")
	}
	groupRes, err := mappers.GroupModelToGroupRes(group, user, currency)
	if err != nil {
		return nil, errors.New("Unable tor process request")
	}
	return groupRes, nil
}
