package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"mini-api/config"
)

type AuthMiddleware struct {
	cfg *config.APIConfig
}

func NewAuthMiddleware(
	cfg *config.APIConfig,
) AuthMiddleware {
	return AuthMiddleware{
		cfg: cfg,
	}
}

func (mw AuthMiddleware) AuthToken() echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == mw.cfg.AuthToken, nil
	})
}
