package userauth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var jwt_secret = []byte("myunsafesecret")

func GenerateToken(userId int) string {

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

func GetUserIdAndValidateToken(tokenString string) (string, error) {

	claims, err := validateSignature(tokenString)
	if err != nil {
		return "", err
	}
	return claims["userId"].(string), nil
}

func validateSignature(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwt_secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
