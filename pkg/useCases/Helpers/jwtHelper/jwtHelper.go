package jwtHelper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtHelper interface {
	GenerateToken(email string, secret []byte, expireTime time.Duration) (string, error)
	ValidateToken(token string, secret []byte) (string, error)
}

type JwtHelperImpl struct {
}

func (j *JwtHelperImpl) GenerateToken(email string, secret []byte, expireTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(expireTime).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JwtHelperImpl) ValidateToken(token string, secret []byte) (string, error) {
	claims := jwt.MapClaims{}

	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return "", err
	}

	if !tk.Valid {
		return "", err
	}

	email := claims["email"].(string)

	return email, nil
}
