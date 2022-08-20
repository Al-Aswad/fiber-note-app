package requests

import "github.com/go-playground/validator/v10"

type CreateTodo struct {
	ActivityGroupID int    `validate:"required" json:"activity_group_id"`
	Title           string `validate:"required" json:"title"`
}

var validateTodo = validator.New()

func ValidateCreateTodo(todo CreateTodo) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validateTodo.Struct(todo)
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
