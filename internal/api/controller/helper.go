package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"mini-api/internal/api/constant"
	"mini-api/internal/api/dto"
	"mini-api/lib"
)

func GetSession(c echo.Context) dto.Session {
	if session, ok := c.Get(constant.KeySession).(dto.Session); ok {
		return session
	}
	return dto.Session{}
}

func NewCommonErrorBadRequest(err error) error {
	return lib.CommonError{
		StatusCode:    http.StatusBadRequest,
		ErrorCode:     constant.ErrCodeBadRequest,
		ErrorInstance: err,
	}
}
