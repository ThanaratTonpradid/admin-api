package middleware

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"mini-api/internal/api/constant"
	"mini-api/internal/api/dto"
	"mini-api/internal/api/service"
	"mini-api/lib"
)

type JWTAuthMiddleware struct {
	jwtHandler  *lib.JWTHandler
	authService service.AuthService
}

func NewJWTAuthMiddleware(
	jwtHandler *lib.JWTHandler,
	authService service.AuthService,
) JWTAuthMiddleware {
	return JWTAuthMiddleware{
		jwtHandler:  jwtHandler,
		authService: authService,
	}
}

func (mw JWTAuthMiddleware) JWTAuth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:         &jwt.StandardClaims{},
		SigningKey:     mw.jwtHandler.Options.JWTSecret,
		ParseTokenFunc: mw.ParseToken,
	})
}

func (mw JWTAuthMiddleware) ParseToken(auth string, c echo.Context) (interface{}, error) {
	claims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(auth, &claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != middleware.AlgorithmHS256 {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return mw.jwtHandler.Options.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	session, err := mw.GetSession(claims)
	if err != nil {
		return nil, err
	}
	c.Set(constant.KeySession, session)
	return token, nil
}

func (mw JWTAuthMiddleware) GetSession(claims jwt.StandardClaims) (dto.Session, error) {
	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return dto.Session{}, err
	}
	session, err := mw.authService.GetSession(uint32(id))
	if err != nil {
		return dto.Session{}, errors.New("session not found")
	}
	if claims.Id != session.ID {
		return dto.Session{}, errors.New("invalid session")
	}
	return session, nil
}
