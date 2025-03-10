package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tejashwinn/splitwise/constants"
	"github.com/tejashwinn/splitwise/mappers"
	"github.com/tejashwinn/splitwise/repos"
	"github.com/tejashwinn/splitwise/types"
	"github.com/tejashwinn/splitwise/utils"
)

type GroupHandler struct {
	UserRepo      repos.UserRepo
	GroupRepo     repos.GroupRepo
	GroupUserRepo repos.GroupUserRepo
	CurrencyRepo  repos.CurrencyRepo
	JwtUtil       utils.JwtUtil
}

func NewGroupHandler(
	userRepo repos.UserRepo,
	groupRepo repos.GroupRepo,
	currencyRepo repos.CurrencyRepo,
	groupUserRepo repos.GroupUserRepo,

	jwtUtil *utils.JwtUtil,
) *GroupHandler {
	return &GroupHandler{UserRepo: userRepo, GroupRepo: groupRepo, CurrencyRepo: currencyRepo, JwtUtil: *jwtUtil, GroupUserRepo: groupUserRepo}
}

func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	req := &types.GroupReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}
	userId, err := h.JwtUtil.GetUserId(r.Context())

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
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
	defaultUser := &types.GroupUser{
		GroupId: group.Id,
		UserId:  userId,
	}
	h.GroupUserRepo.Save(r.Context(), defaultUser)
	groupRes, err := h.groupToGroupRes(group, w, r)
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(groupRes)
}

func (h *GroupHandler) ListGroupUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Get path variables
	groupId, err := strconv.ParseInt(vars["groupId"], 10, 64)

	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid group Id", http.StatusBadRequest)
		return
	}
	usersInGroup, err := h.GroupUserRepo.FindUserIdByGroupId(context.Background(), groupId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to fetch users", http.StatusBadRequest)
		return
	}
	users, err := h.UserRepo.FindByIdIn(context.Background(), usersInGroup)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to fetch users", http.StatusBadRequest)
		return
	}
	usersRes, err := mappers.MapUsersToUserRes(users)

	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(usersRes)
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
