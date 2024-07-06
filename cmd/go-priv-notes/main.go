package main

import (
	"log"
	"net"
	"net/http"

	"github.com/samocodes/go-priv-notes/db"
	"github.com/samocodes/go-priv-notes/env"
	"github.com/samocodes/go-priv-notes/internal/api"
)

func init() {
	env.Load()
	db.Load()
}

type Application struct {
	config env.Config
}

func (app *Application) Serve() error {
	srv := &http.Server{
		Addr:    net.JoinHostPort("localhost", app.config.PORT),
		Handler: api.Router(),
	}

	log.Printf("🚀 Server listening to port %s", app.config.PORT)

	return srv.ListenAndServe()
}

func main() {
	app := Application{
		config: env.DefaultConfig,
	}

	if err := app.Serve(); err != nil {
		log.Fatalf("Error starting server: %s\n", err.Error())
	}
}
