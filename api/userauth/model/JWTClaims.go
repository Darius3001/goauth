package userauth

import "time"

type JWTClaims struct {
	UserId         int
	ExpirationDate time.Time
}
