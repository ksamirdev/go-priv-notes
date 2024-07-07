package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router(db *sql.DB) http.Handler {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/user", UserRouter(db))
		r.Mount("/note", NoteRouter(db))
	})

	return router
}
