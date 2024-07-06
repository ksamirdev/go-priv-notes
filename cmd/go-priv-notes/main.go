package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samocodes/go-priv-notes/env"
	"github.com/samocodes/go-priv-notes/internal/api"
)

func init() {
	env.Load()
}

type Application struct {
	config env.Config
}

func (app *Application) Serve() error {
	port := app.config.PORT

	log.Printf("🚀 Server listening to port %s", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: api.Router(),
	}

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
