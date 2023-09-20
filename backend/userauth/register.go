package userauth

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "golang.org/x/crypto/bcrypt"
)

type registrationRequest struct {
	Name 			string `json:"username"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

func handleRegistration(w http.ResponseWriter, r *http.Request) {

	var registrationRequest registrationRequest

	decoder := json.NewDecoder(r.Body)	

	err := decoder.Decode(&registrationRequest)
	if err != nil {
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}


	fmt.Println(registrationRequest)
	w.WriteHeader(http.StatusOK)
}