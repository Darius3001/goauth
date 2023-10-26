package userauth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"openpager.com/m/database"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var loginRequest loginRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&loginRequest)
	if err != nil {
		http.Error(w, "failed to Decode JSON body", http.StatusBadRequest)
		return
	}

	var password_hash []byte

	err = database.ExecuteQuery(func(db *sql.DB) error {
		return db.
			QueryRow(`SELECT password_hash FROM users`).
			Scan(&password_hash)
	})

	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}

	err = bcrypt.CompareHashAndPassword(password_hash, []byte(loginRequest.Password))

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
