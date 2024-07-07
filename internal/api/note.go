package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/samocodes/go-priv-notes/helpers"
	"github.com/samocodes/go-priv-notes/types"
)

func NoteRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	r.Post("/send", func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsURLEncodedFormValid(r) {
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
			http.Error(w, helpers.ErrInvalidUsername, http.StatusBadRequest)
			return
		}

		var user types.UsersTable
		row := db.QueryRow("SELECT * FROM users WHERE username = ?", username)
		if err := row.Scan(&user.Username, &user.Pin, &user.CreatedAt); err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "User with this username does not exist", http.StatusBadRequest)
				return
			}

			http.Error(w, "Error querying user", http.StatusInternalServerError)
			return
		}

		id, err := uuid.NewV7()
		if err != nil {
			http.Error(w, fmt.Sprintf("UUIDV7 raised an issue: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		_, err = db.Exec(`INSERT INTO notes(id, content, username) VALUES (?, ?, ?)`, id.String(), content, username)
		if err != nil {
			http.Error(w, "Issue when inserting the item", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Sent the note!")
	})

	return r
}
