package gojwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// VARIABLES

var secretKey = []byte("ca01a4bf0291a4d435323e464c90505e7ef65828c146ddbde6d01d449ad99c6a")

// JWT

func ConfigJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"sub": username,
	})

	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func VerifyJWT(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return jwt.ErrTokenUnverifiable
	}
	return nil
}
