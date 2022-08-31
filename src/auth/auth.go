package auth

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateJSONWebToken(id uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}
