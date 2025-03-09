package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/tejashwinn/splitwise/constants"
	"github.com/tejashwinn/splitwise/repositories"
	"github.com/tejashwinn/splitwise/types"
)

type UserHandler struct {
	Repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
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
	err = validateReq(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err = h.Repo.InsertOneUser(context.Background(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	json.NewEncoder(w).Encode(user)
}

func validateReq(req *types.User) error {
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
