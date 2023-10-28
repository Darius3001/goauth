package userauth

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	userauth "openpager.com/m/userauth/model"
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

func ValidateToken(tokenString string) (*userauth.JWTClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwt_secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("Extracting token claims failed")
	}

	userIdFloat := claims["userId"].(float64)
	expirationTimeStamp := int64(claims["expirationDate"].(float64))

	result := &userauth.JWTClaims{
		UserId:         int(userIdFloat),
		ExpirationDate: time.Unix(expirationTimeStamp, 0),
	}

	return result, nil
}
