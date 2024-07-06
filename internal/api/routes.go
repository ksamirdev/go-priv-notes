package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", userResource{}.Routes())
		r.Mount("/notes", notesResources{}.Routes())
	})

	return router
}
