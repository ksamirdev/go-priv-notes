package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/samocodes/go-priv-notes/db"
	"github.com/samocodes/go-priv-notes/helpers"
)

type notesResources struct{}

func (nr notesResources) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/send", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		content := r.FormValue("content")
		username := r.FormValue("username")

		if content == "" || username == "" {
			http.Error(w, "Either content or username was not provided", http.StatusBadRequest)
			return
		}

		if !helpers.IsValidUsername(username) {
			http.Error(w, "Username is not valid", http.StatusBadRequest)
			return
		}

		id, err := uuid.NewV7()
		if err != nil {
			http.Error(w, "UUIDV7 raised an issue", http.StatusInternalServerError)
			return
		}

		_, err = db.DB.Exec(`INSERT INTO notes(id, content, username) VALUES (?, ?, ?)`, id.String(), content, username)
		if err != nil {
			http.Error(w, "Issue when inserting the item", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Sent the note!")
	})

	return r
}
