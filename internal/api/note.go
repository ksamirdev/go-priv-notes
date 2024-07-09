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
			fmt.Fprintf(w, "Invalid request")
			return
		}

		content := r.FormValue("content")
		username := r.FormValue("username")

		if content == "" || username == "" {
			fmt.Fprintf(w, "Either content or username was not provided")
			return
		}

		if !helpers.IsValidUsername(username) {
			fmt.Fprintf(w, helpers.ErrInvalidUsername)
			return
		}

		var user types.UsersTable
		row := db.QueryRow("SELECT * FROM users WHERE username = ?", username)
		if err := row.Scan(&user.Username, &user.Pin, &user.CreatedAt); err != nil {
			if err == sql.ErrNoRows {
				fmt.Fprintf(w, "User with this username does not exist")
				return
			}

			fmt.Fprintf(w, "Error querying user")
			return
		}

		id, err := uuid.NewV7()
		if err != nil {
			fmt.Fprintf(w, "UUIDV7 raised an issue: %s", err.Error())
			return
		}

		_, err = db.Exec(`INSERT INTO notes(id, content, username) VALUES (?, ?, ?)`, id.String(), content, username)
		if err != nil {
			fmt.Fprintf(w, "Issue when inserting the item")
			return
		}

		fmt.Fprint(w, "Sent the note!")
	})

	return r
}
