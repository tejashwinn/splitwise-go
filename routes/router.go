package routes

import (
	"github.com/gorilla/mux"
	"github.com/tejashwinn/splitwise/handlers"
	"github.com/tejashwinn/splitwise/middleware"
)

func SetupRouter(
	auth *middleware.AuthMiddleware,
	userH *handlers.UserHandler,
	currencyH *handlers.CurrencyHandler,
	groupH *handlers.GroupHandler,
) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/users/signup", userH.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users/login", userH.Login).Methods("POST")
	r.HandleFunc("/api/v1/health", handlers.Health).Methods("GET")
	r.HandleFunc("/api/v1/ping", handlers.Ping).Methods("GET")

	protected := r.PathPrefix("/api/v1").Subrouter()
	protected.Use(auth.AuthenticateAndSetUserId)

	protected.HandleFunc("/users/whoami", userH.WhoAmI).Methods("GET")
	protected.HandleFunc("/currencies", currencyH.ListCurrencies).Methods("GET")
	protected.HandleFunc("/groups", groupH.CreateGroup).Methods("POST")
	return r
}
