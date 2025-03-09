package handlers

import (
	"context"
	"encoding/json"

	"log"
	"net/http"

	"github.com/tejashwinn/splitwise/constants"
	"github.com/tejashwinn/splitwise/mappers"
	"github.com/tejashwinn/splitwise/repositories"
	"github.com/tejashwinn/splitwise/types"
	"github.com/tejashwinn/splitwise/util"
)

type UserHandler struct {
	Repo    repositories.UserRepo
	JwtUtil util.JwtUtil
}

func NewUserHandler(
	repo repositories.UserRepo,
	jwtUtil *util.JwtUtil,
) *UserHandler {
	return &UserHandler{Repo: repo, JwtUtil: *jwtUtil}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.GetAllUsers(context.Background())
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	usersRes := []types.UserRes{}
	for _, user := range users {
		userRes, err := mappers.MapUserToUserRe(
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

	user, err = h.Repo.InsertOneUser(context.Background(), user)
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

	user, err := h.Repo.FindByEmailOrUsername(context.Background(), req.User)
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
