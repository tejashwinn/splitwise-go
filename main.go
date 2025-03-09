package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tejashwinn/splitwise/configs"
	"github.com/tejashwinn/splitwise/handlers"
	"github.com/tejashwinn/splitwise/middleware"
	"github.com/tejashwinn/splitwise/repos"
	"github.com/tejashwinn/splitwise/routes"
	"github.com/tejashwinn/splitwise/types"
	"github.com/tejashwinn/splitwise/utils"
	"go.uber.org/fx"
)

// StartServer starts the HTTP server with lifecycle hooks
func StartServer(
	lc fx.Lifecycle,
	cfg *types.Config,
	router *mux.Router,
) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Starting server on %s\n", server.Addr)
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Printf("Shutting down server on %s\n", server.Addr)
			return server.Close()
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			configs.LoadConfig,
			configs.ConnectDB,

			repos.NewUserRepository,
			repos.NewCurrencyRepository,
			repos.NewGroupRepository,
			repos.NewGroupUserMapRepository,

			middleware.NewAuthMiddleware,

			utils.NewJwtUtil,

			handlers.NewUserHandler,
			handlers.NewCurrencyHandler,
			handlers.NewGroupHandler,

			routes.SetupRouter,
		),
		fx.Invoke(StartServer),
	)

	app.Run()
}
