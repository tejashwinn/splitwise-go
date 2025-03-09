package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/tejashwinn/splitwise/repositories"
)

// UserHandler struct (depends on UserRepository)
type UserHandler struct {
	Repo repositories.UserRepository
}

// NewUserHandler initializes UserHandler with a repository
func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

// GetUsers handles GET /users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.GetAllUsers(context.Background())
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
