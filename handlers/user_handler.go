package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"log"
	"net/http"

	"github.com/tejashwinn/splitwise/constants"
	"github.com/tejashwinn/splitwise/mappers"
	repositories "github.com/tejashwinn/splitwise/repos"
	"github.com/tejashwinn/splitwise/types"
	"github.com/tejashwinn/splitwise/util"
)

type UserHandler struct {
	UserRepo repositories.UserRepo
	JwtUtil  util.JwtUtil
}

func NewUserHandler(
	userRepo repositories.UserRepo,
	jwtUtil *util.JwtUtil,
) *UserHandler {
	return &UserHandler{UserRepo: userRepo, JwtUtil: *jwtUtil}
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserRepo.FindAll(context.Background())
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	usersRes := []types.UserRes{}
	for _, user := range users {
		userRes, err := mappers.MapUserToUserRes(
			&user,
		)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		}
		usersRes = append(usersRes, *userRes)
	}
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(usersRes)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req types.UserReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}
	err = util.ValidateCreateUser(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.Password, err = util.HashPassword(req.Password)
	if err != nil {
		log.Printf("Unable to hash password: %s with error: %s",
			req.Password,
			err.Error(),
		)
		http.Error(w, "Unable to create user", http.StatusConflict)
		return
	}
	user, err := mappers.CreateReqToModel(&req)

	user, err = h.UserRepo.Save(context.Background(), user)
	if err != nil {
		http.Error(w, "Unable to  create user", http.StatusConflict)
		return
	}

	accessToken, refreshToken, err := h.JwtUtil.GenerateToken(user)
	if err != nil {
		http.Error(w, "Unable to generate token", http.StatusConflict)
		return
	}
	token := &types.TokenRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(token)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req types.LoginReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}
	err = util.ValidateLoginUser(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.UserRepo.FindByEmailOrUsername(context.Background(), req.User)
	if err != nil || !util.CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	accessToken, refreshToken, err := h.JwtUtil.GenerateToken(user)
	if err != nil {
		http.Error(w, "Unable to generate token", http.StatusConflict)
		return
	}
	token := &types.TokenRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(token)
}

func (h *UserHandler) WhoAmI(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(
		fmt.Sprint(r.Context().Value(constants.UserId)),
		10,
		64,
	)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	user, err := h.UserRepo.FindById(context.Background(), userId)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	userRes, err := mappers.MapUserToUserRes(user)
	if err != nil {
		http.Error(w, "unable to map user", http.StatusInternalServerError)
		return
	}
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(userRes)
}
