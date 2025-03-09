package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/tejashwinn/splitwise/configurations"
	"github.com/tejashwinn/splitwise/repositories"
	"github.com/tejashwinn/splitwise/routes"
	"github.com/tejashwinn/splitwise/types"
	"go.uber.org/fx"
)

// StartServer starts the HTTP server with lifecycle hooks
func StartServer(
	lc fx.Lifecycle,
	cfg *types.Config,
	router *mux.Router,

) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: router,
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
			config.LoadConfig,              // Load config
			config.ConnectDB,               // Connect to DB
			repositories.NewUserRepository, // Inject UserRepository
			routes.SetupRouter,             // Inject Router
		),
		fx.Invoke(StartServer),
	)

	app.Run()
}
