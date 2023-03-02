package web

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenProvider struct {
	secret   string
	validity uint
	issuer   string
}

func NewTokenProvider(secret string, validity int64, issuer string) *TokenProvider {
	return &TokenProvider{secret: secret, validity: uint(validity), issuer: issuer}
}

func (p *TokenProvider) GetToken(userName string, currentTime time.Time) (string, error) {
	exp := currentTime.Add(time.Minute * time.Duration(p.validity))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    p.issuer,
		Subject:   userName,
		ExpiresAt: jwt.NewNumericDate(exp),
	})
	tokenString, err := token.SignedString([]byte(p.secret))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
