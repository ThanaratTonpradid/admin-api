package lib

type CommonError struct {
	StatusCode    int
	ErrorCode     string
	ErrorInstance error
}

type ValidationError struct {
	ErrorMessage string
	ErrorDetail  []ValidationErrorDetail
}

type ValidationErrorDetail struct {
	Field   string `json:"field" example:"id"`
	Tag     string `json:"tag" example:"required"`
	Message string `json:"message" example:"Key: 'Member.id' Error:Field validation for 'id' failed on the 'required' tag"`
}

func (e CommonError) Error() string {
	return e.ErrorInstance.Error()
}

func (e ValidationError) Error() string {
	return e.ErrorMessage
}
