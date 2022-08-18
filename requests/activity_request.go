package requests

import "github.com/go-playground/validator/v10"

type CreateActivity struct {
	Title string `validate:"required" json:"title"`
	Email string `validate:"required" json:"email"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateCreateActivity(activity CreateActivity) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(activity)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
