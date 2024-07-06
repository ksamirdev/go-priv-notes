package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/samocodes/go-priv-notes/crypto"
	"github.com/samocodes/go-priv-notes/db"
	"github.com/samocodes/go-priv-notes/helpers"
)

type userResource struct{}

func (rs userResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsURLEncodedFormValid(r) {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		pin := r.FormValue("pin")

		err := rs.createUser(username, pin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}

		fmt.Fprintln(w, "User has been created!")
	})

	return r
}

func (rs userResource) createUser(username, pin string) error {
	if !helpers.IsValidUsername(username) {
		return fmt.Errorf("invalid username provided")
	}

	if !helpers.IsValidPin(pin) {
		return fmt.Errorf("invalid pin provided")
	}

	hashedPin, err := crypto.AESEncrypt(pin)
	if err != nil {
		return err
	}

	if _, err := db.DB.Exec(`INSERT INTO users(username, pin) VALUES (?, ?)`, username, hashedPin); err != nil {
		return err
	}

	return nil
}

func (rs userResource) accessUser(username, pin string) {
}
