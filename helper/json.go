package helper

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ParseJSON(body string, result interface{}) error {
	if err := json.Unmarshal([]byte(body), result); err != nil {
		return err
	}
	return validate.Struct(result)
}

func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}
