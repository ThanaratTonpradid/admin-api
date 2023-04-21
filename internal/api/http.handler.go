package api

import (
	"fmt"
	"net/http"

	"github.com/dollarsignteam/go-logger"
	"github.com/labstack/echo/v4"

	"mini-api/internal/api/constant"
	"mini-api/internal/api/dto"
	"mini-api/lib"
)

func customHTTPErrorHandler(log *logger.Logger) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		if commonError, ok := err.(lib.CommonError); ok {
			errResp := dto.ErrorResponse{
				Success:    false,
				StatusCode: commonError.StatusCode,
				ErrorCode:  commonError.ErrorCode,
			}
			if commonError.ErrorInstance != nil {
				errResp.ErrorDescription = commonError.ErrorInstance.Error()
			}
			if validationError, ok := commonError.ErrorInstance.(lib.ValidationError); ok {
				errorValidationDetail := []dto.ErrorValidationDetail{}
				for _, e := range validationError.ErrorDetail {
					errorValidationDetail = append(errorValidationDetail, dto.ErrorValidationDetail{
						ValidationErrorDetail: e,
					})
				}
				errValidationResp := dto.ErrorValidationResponse{
					ErrorResponse:   errResp,
					ErrorValidation: errorValidationDetail,
				}
				if err := errValidationResp.JSON(c); err != nil {
					log.Error(err)
				}
				return
			}
			if err := errResp.JSON(c); err != nil {
				log.Error(err)
			}
			return
		}
		if ok := httpErrorHandler(err, c, log); ok {
			return
		}
		resp := lib.NewErrorResponse(err)
		if err := resp.JSON(c); err != nil {
			log.Error(err)
		}
	}
}

func httpErrorHandler(err error, c echo.Context, log *logger.Logger) bool {
	if errHttp, ok := err.(*echo.HTTPError); ok {
		code := constant.ErrCodeSomethingWentWrong
		errResp := dto.ErrorResponse{
			Success:          false,
			StatusCode:       errHttp.Code,
			ErrorDescription: fmt.Sprintf("%v", errHttp.Message),
		}
		switch errHttp.Code {
		case http.StatusBadRequest:
			code = constant.ErrCodeBadRequest
		case http.StatusUnauthorized:
			code = constant.ErrCodeUnauthorized
		}
		errResp.ErrorCode = code
		if err := errResp.JSON(c); err != nil {
			log.Error(err)
		}
		return true
	}
	return false
}
