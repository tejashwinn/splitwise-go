package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tejashwinn/splitwise/handlers"
	"github.com/tejashwinn/splitwise/middlewares"
)

func SetupRouter(
	authMiddleware *middlewares.AuthMiddleware,
	userHandler *handlers.UserHandler,
) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users/signup", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users/login", userHandler.Login).Methods("POST")
	r.Handle(
		"/api/v1/users/whoami",
		authMiddleware.AuthenticateAndSetUserId(
			http.HandlerFunc(userHandler.WhoAmI),
		),
	).Methods("GET")

	r.HandleFunc("/api/v1/health", handlers.Health).Methods("GET")
	r.HandleFunc("/api/v1/ping", handlers.Ping).Methods("GET")
	return r
}
