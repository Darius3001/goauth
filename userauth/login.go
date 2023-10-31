package userauth

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"goauth.com/m/database"
	userauth "goauth.com/m/userauth/model"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var loginRequest userauth.LoginRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&loginRequest)
	if err != nil {
		http.Error(w, "failed to Decode JSON body", http.StatusBadRequest)
		return
	}

	var password_hash []byte
	var userId int

	err = database.ExecuteQuery(func(db *sql.DB) error {
		return db.
			QueryRow(`SELECT id, password_hash FROM users`).
			Scan(&userId, &password_hash)
	})

	if err == sql.ErrNoRows {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword(password_hash, []byte(loginRequest.Password))

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	response := userauth.LoginResponse{
		Token: GenerateToken(userId),
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Panicf("Error encoding JSON in login. %s", err)
		http.Error(w, "Internal encoding error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(responseJson)
}
