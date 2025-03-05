package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
)

var signingKey = os.Getenv("SIGN_KEY_STRING")

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func CreateToken(userId int) (token string) {
	params := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		UserId: userId,
	})

	token, _ = params.SignedString([]byte(signingKey))
	return token
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	return claims.UserId, nil
}
