package lib

import (
	"time"

	"github.com/golang-jwt/jwt"

	"mini-api/helper"
)

type JWTToken struct {
	ID        string
	Token     string
	IssuedAt  int64
	ExpiresAt int64
}

type JWTOptions struct {
	JWTSecret     []byte
	JWTExpiresTTL time.Duration
}

type JWTHandler struct {
	Options JWTOptions
}

func NewJWTHandler(opts JWTOptions) *JWTHandler {
	return &JWTHandler{
		Options: opts,
	}
}

func (h JWTHandler) CreateToken(subject string) (JWTToken, error) {
	id := helper.UUID()
	issuedAt := time.Now().Unix()
	expiresAt := time.Now().Add(h.Options.JWTExpiresTTL).Unix()
	claims := &jwt.StandardClaims{
		Id:        id,
		ExpiresAt: expiresAt,
		IssuedAt:  issuedAt,
		Subject:   subject,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.Options.JWTSecret)
	return JWTToken{
		ID:        id,
		Token:     tokenString,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}, err
}
