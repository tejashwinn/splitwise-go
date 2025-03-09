package routes

import (
	"github.com/gorilla/mux"
	"github.com/tejashwinn/splitwise/handlers"
	"github.com/tejashwinn/splitwise/repositories"
)

// SetupRouter initializes the router with handlers
func SetupRouter(repo repositories.UserRepository) *mux.Router {
	r := mux.NewRouter()

	userHandler := handlers.NewUserHandler(repo)
	r.HandleFunc("/api/v1/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/health", handlers.Health).Methods("GET")
	r.HandleFunc("/api/v1/ping", handlers.Ping).Methods("GET")
	return r
}
