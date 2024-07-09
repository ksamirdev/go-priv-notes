package api

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/samocodes/go-priv-notes/crypto"
	"github.com/samocodes/go-priv-notes/helpers"
	"github.com/samocodes/go-priv-notes/types"
)

type NotesResponse struct {
	Error   bool
	Message string

	Username string
	Notes    []types.Notes
}

func UserRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	// TODO: pagination
	r.Get("/notes", func(w http.ResponseWriter, r *http.Request) {
		// gets username, and pin from url query
		username := r.URL.Query().Get("username")
		pin := r.URL.Query().Get("pin")

		tmpl := template.Must(template.ParseFiles("internal/template/notes.html"))

		if !helpers.IsValidUsername(username) || !helpers.IsValidPin(pin) {
			tmpl.Execute(w, NotesResponse{Error: true, Message: "Either username or pin is invalid", Username: username})
			return
		}

		// check if user exists or create one
		if err := createOrFindUser(username, pin, db); err != nil {
			tmpl.Execute(w, NotesResponse{Error: true, Message: err.Error(), Username: username})
			return
		}

		// fetch user's notes
		notes, err := fetchUserNotes(username, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tmpl.Execute(w, NotesResponse{Username: username, Notes: notes, Error: false})
	})

	return r
}

func createUser(username, pin string, db *sql.DB) error {
	hashedPin, err := crypto.Encrypt(pin)
	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO users(username, pin) VALUES (?, ?)`, username, hashedPin)
	if err != nil {
		return err
	}

	return nil
}

func createOrFindUser(username, pin string, db *sql.DB) error {
	// fetch the user
	var user types.UsersTable

	query := "SELECT username, pin, created_at FROM users WHERE username = ?"
	row := db.QueryRow(query, username)
	err := row.Scan(&user.Username, &user.Pin, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// if doesnt exists, create one
			if err := createUser(username, pin, db); err != nil {
				return err
			}

			return nil
		}

		return fmt.Errorf("Invalid user credentials")
	}

	p, err := crypto.Decrypt(user.Pin)
	if err != nil || p != pin {
		return fmt.Errorf("Invalid user credentials")
	}

	return nil
}

func fetchUserNotes(username string, db *sql.DB) ([]types.Notes, error) {
	var notes []types.Notes

	query := "SELECT id, content, created_at FROM notes WHERE username = ? ORDER BY created_at DESC"
	rows, err := db.Query(query, username)
	if err != nil {
		return notes, errors.New("error while fetching user's notes")
	}
	defer rows.Close()

	for rows.Next() {
		var note types.NotesTable
		if err := rows.Scan(&note.Id, &note.Content, &note.CreatedAt); err != nil {
			continue
		}

		notes = append(notes, types.Notes{
			Id:        note.Id,
			Content:   note.Content,
			CreatedAt: helpers.ReadableTime(note.CreatedAt),
		})
	}

	return notes, nil
}
