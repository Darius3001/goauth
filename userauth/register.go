package userauth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"goauth.com/m/database"
	userauth "goauth.com/m/userauth/model"
	"golang.org/x/crypto/bcrypt"
)

func handleRegistration(w http.ResponseWriter, r *http.Request) {

	var registrationRequest userauth.RegistrationRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&registrationRequest)
	if err != nil {
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(registrationRequest.Password),
		bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	//TODO email verification
	//TODO better error on duplicate email
	err = database.ExecuteQuery(func(db *sql.DB) error {
		_, err := db.Query(fmt.Sprintf(`
			INSERT INTO users (username, email, password_hash)
			VALUES ('%s', '%s', '%s')
		`, registrationRequest.Name, registrationRequest.Email, hashedPassword))
		return err
	})

	if err != nil {
		http.Error(w, "Error inserting to database", http.StatusInternalServerError)
		fmt.Println("WARNING: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
