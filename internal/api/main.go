package api

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router(db *sql.DB) http.Handler {
	router := chi.NewRouter()

	// serve css
	router.Get("/dist/main.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "internal/template/dist/main.css")
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("internal/template/index.html"))
		tmpl.Execute(w, nil)
	})

	router.Mount("/user", UserRouter(db))
	router.Mount("/note", NoteRouter(db))

	return router
}
