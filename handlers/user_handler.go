package handlers

import (
	"context"
	"encoding/json"

	"log"
	"net/http"

	"github.com/tejashwinn/splitwise/constants"
	"github.com/tejashwinn/splitwise/repositories"
	"github.com/tejashwinn/splitwise/types"
	"github.com/tejashwinn/splitwise/util"
)

type UserHandler struct {
	Repo    repositories.UserRepository
	JwtUtil util.JwtUtil
}

func NewUserHandler(
	repo repositories.UserRepository,
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
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}
	err = util.ValidateCreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		log.Printf("Unable to hash password: %s with error: %s",
			user.Password,
			err.Error(),
		)
		http.Error(w, "Unable to create user", http.StatusConflict)
		return
	}

	user, err = h.Repo.InsertOneUser(context.Background(), &user)
	if err != nil {
		http.Error(w, "Unable to  create user", http.StatusConflict)
		return
	}

	accessToken, refreshToken, err := h.JwtUtil.GenerateToken(&user)
	if err != nil {
		http.Error(w, "Unable to generate token", http.StatusConflict)
		return
	}
	token := &types.TokenResposne{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(token)
}
