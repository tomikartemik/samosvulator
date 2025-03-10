package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var signingKey = os.Getenv("SIGN_KEY_STRING")

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func CreateToken(userId int) (string, error) {
	if userId == 0 {
		return "", errors.New("user_id cannot be empty")
	}

	params := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
		UserId: userId,
	})

	token, err := params.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return token, nil
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

	if claims.UserId == 0 {
		return 0, errors.New("user_id is empty or invalid")
	}

	return claims.UserId, nil
}
