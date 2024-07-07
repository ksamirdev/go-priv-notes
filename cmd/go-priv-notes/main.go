package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/samocodes/go-priv-notes/db"
	"github.com/samocodes/go-priv-notes/env"
	"github.com/samocodes/go-priv-notes/internal/api"
)

func init() {
	env.Load()
}

type Application struct {
	config env.Config
	db     *sql.DB
}

func (app *Application) Serve() error {
	srv := &http.Server{
		Addr:    net.JoinHostPort("localhost", app.config.PORT),
		Handler: api.Router(app.db),
	}

	log.Printf("[api] 🚀 listening to port %s", app.config.PORT)

	return srv.ListenAndServe()
}

func main() {
	db := db.Load()
	defer db.Close()

	app := Application{
		config: env.DefaultConfig,
		db:     db,
	}

	if err := app.Serve(); err != nil {
		log.Fatalf("[api] Error starting: %s\n", err.Error())
	}
}
