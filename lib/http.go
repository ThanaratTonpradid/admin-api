package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/dollarsignteam/go-logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Validator struct {
	validate *validator.Validate
}

type BinderWithValidation struct{}

type HttpHandler struct {
	Engine   *echo.Echo
	Validate *validator.Validate
}

func NewHttpHandler() *HttpHandler {
	engine := echo.New()
	engine.HidePort = true
	engine.HideBanner = true
	engine.Binder = &BinderWithValidation{}
	engine.Pre(middleware.RemoveTrailingSlash())
	engine.Use(middleware.CORS())
	engine.GET("/", func(c echo.Context) error {
		return Response{
			Message: "200 OK",
		}.JSON(c)
	})
	httpHandler := &HttpHandler{
		Engine: engine,
	}
	httpHandler.Engine.HTTPErrorHandler = customHTTPErrorHandler(log)
	httpHandler.Engine.Validator = func() echo.Validator {
		v := validator.New()
		if err := v.RegisterValidation("json", func(fl validator.FieldLevel) bool {
			var js json.RawMessage
			return json.Unmarshal([]byte(fl.Field().String()), &js) == nil
		}); err != nil {
			log.Warnf("Register validation json: %s", err.Error())
		}
		if err := v.RegisterValidation("in", func(fl validator.FieldLevel) bool {
			value := fl.Field().String()
			return (containsString(strings.Split(fl.Param(), ";"), value) || value == "")
		}); err != nil {
			log.Warnf("Register validation in: %s", err.Error())
		}
		return &Validator{validate: v}
	}()
	return httpHandler
}

func customHTTPErrorHandler(log *logger.Logger) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		resp := NewErrorResponse(err)
		if err := resp.JSON(c); err != nil {
			log.Error(err)
		}
	}
}

func containsString(items []string, item string) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

func (a *Validator) Validate(i interface{}) error {
	return a.validate.Struct(i)
}

func (BinderWithValidation) Bind(i interface{}, c echo.Context) error {
	binder := &echo.DefaultBinder{}
	if err := binder.Bind(i, c); err != nil {
		return errors.New(err.(*echo.HTTPError).Message.(string))
	}
	if err := c.Validate(i); err != nil {
		var fieldList []string
		var errDetailList []ValidationErrorDetail
		if ves, ok := err.(validator.ValidationErrors); ok {
			for _, ve := range ves {
				fieldList = append(fieldList, fmt.Sprintf("'%s'", ve.Field()))
				errDetailList = append(errDetailList, ValidationErrorDetail{
					Field:   ve.Field(),
					Tag:     ve.Tag(),
					Message: ve.Error(),
				})
			}
			return ValidationError{
				ErrorMessage: fmt.Sprintf("Validation failed for %s", strings.Join(fieldList, ", ")),
				ErrorDetail:  errDetailList,
			}
		}
		return err
	}
	return nil
}
