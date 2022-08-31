package auth

import (
	"api/src/config"
	"errors"
	"net/http"
	"strings"
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

func ValidateJSONWebToken(r *http.Request) error {
	authorization := r.Header.Get("Authorization")

	if len(strings.Split(authorization, " ")) != 2 {
		return errors.New("Malformed token")
	}

	tokenString := strings.Split(authorization, " ")[1]

	token, err := jwt.Parse(tokenString, GetVerifySecret)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Could not validate JWT token")
}

func GetVerifySecret(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("JWT Token has an invalid signature method")
	}

	return config.SecretKey, nil
}
