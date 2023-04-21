package dto

import (
	"github.com/labstack/echo/v4"

	"mini-api/lib"
)

type IErrorResponse interface {
	Error() string
	JSON(c echo.Context) error
}

type ErrorResponse struct {
	StatusCode       int    `json:"-"`
	Success          bool   `json:"success" example:"false"`
	ErrorCode        string `json:"errorCode,omitempty" example:"SOMETHING_WENT_WRONG"`
	ErrorDescription string `json:"errorDescription,omitempty" example:"Internal server error"`
}

type ErrorValidationResponse struct {
	ErrorResponse
	ErrorValidation []ErrorValidationDetail `json:"errorValidation"`
}

type ErrorValidationDetail struct {
	lib.ValidationErrorDetail
}

func (errResp ErrorResponse) JSON(c echo.Context) error {
	return c.JSON(errResp.StatusCode, errResp)
}

func (errResp ErrorResponse) Error() string {
	return errResp.ErrorCode
}

func (errResp ErrorValidationResponse) JSON(c echo.Context) error {
	return c.JSON(errResp.StatusCode, errResp)
}
