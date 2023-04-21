package lib

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var MessageOK = "OK"
var ErrorPrefix = "E"

type Response struct {
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty" example:"OK"`
}

type ErrorResponse struct {
	StatusCode   int         `json:"-"`
	Success      bool        `json:"success" example:"false"`
	ErrorCode    string      `json:"errorCode,omitempty" example:"E000"`
	ErrorMessage interface{} `json:"errorMessage,omitempty"`
}

func NewErrorResponse(err error) ErrorResponse {
	if err == nil {
		return ErrorResponse{}
	}
	if errResp, ok := err.(ErrorResponse); ok {
		return errResp
	}
	return ErrorResponse{
		ErrorMessage: err,
	}
}

func (errResp *ErrorResponse) ParseErrorResponse() {
	switch err := errResp.ErrorMessage.(type) {
	case *echo.HTTPError:
		errResp.StatusCode = err.Code
		errResp.ErrorMessage = err.Message
		if err.Internal != nil {
			errResp.ErrorMessage = fmt.Sprintf("%s - %s", err.Message, err.Internal)
		}
	case error:
		errResp.ErrorMessage = err.Error()
	case nil, string:
		if err == nil || err == "" {
			errResp.ErrorMessage = http.StatusText(http.StatusInternalServerError)
		}
	}
	if errResp.StatusCode == 0 {
		errResp.StatusCode = http.StatusInternalServerError
	}
	if errResp.ErrorCode == "" {
		errResp.ErrorCode = fmt.Sprintf("%s%d", ErrorPrefix, errResp.StatusCode)
	}
}

func (errResp ErrorResponse) JSON(c echo.Context) error {
	errResp.ParseErrorResponse()
	return c.JSON(errResp.StatusCode, errResp)
}

func (errResp ErrorResponse) Error() string {
	errResp.ParseErrorResponse()
	return fmt.Sprintf("%s(%d): %s", errResp.ErrorCode, errResp.StatusCode, errResp.ErrorMessage)
}

func (resp *Response) ParseResponse() {
	if resp.StatusCode == 0 {
		resp.StatusCode = http.StatusOK
	}
	if resp.Message == "" && resp.Data == nil {
		resp.Message = MessageOK
	}
}

func (resp Response) JSON(c echo.Context) error {
	resp.ParseResponse()
	return c.JSON(resp.StatusCode, resp)
}
