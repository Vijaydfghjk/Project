package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

func Generatetoken(userid int) (string, error) {

	claims := jwt.MapClaims{

		"user_id": userid,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	signedtoken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return signedtoken, nil
}
