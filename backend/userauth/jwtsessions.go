package userauth

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwt_secret []byte = []byte("myunsafesecret")

func generateToken(userId int) string {

	validityDuration := time.Hour * 4

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = jwt.MapClaims{
		"userId":         userId,
		"expirationDate": time.Now().Add(validityDuration).Unix(),
	}

	signedToken, err := token.SignedString(jwt_secret)

	if err != nil {
		log.Fatalf("Error on generating jwt. %s", err)
	}

	return signedToken
}
