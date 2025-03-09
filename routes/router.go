package routes

import (
	"github.com/gorilla/mux"
	"github.com/tejashwinn/splitwise/handlers"
)

// SetupRouter initializes the router with handlers
func SetupRouter(userHandler *handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/health", handlers.Health).Methods("GET")
	r.HandleFunc("/api/v1/ping", handlers.Ping).Methods("GET")
	return r
}
