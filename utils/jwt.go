package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SecretKey string = "asdxcsddasd123*&sd"

type LoginClaims struct {
	Uid int `json:"uid"`
	jwt.StandardClaims
}

func GenerateJwt(uid int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &LoginClaims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	})
	tokens, err := token.SignedString([]byte("123"))
	return tokens, err
}

func ParseToken(tokenString string) (*LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return LoginClaims{}, nil
	})
	if err != nil {
		return nil, err
	}
	if LoginClaims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
		return LoginClaims, nil
	}
	return nil, errors.New("invalid token")
}
