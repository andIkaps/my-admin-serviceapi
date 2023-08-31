package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Tag     string `json:"tag"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

func formatValidationErrors(errs validator.ValidationErrors) []ErrorResponse {
	var errors []ErrorResponse

	for _, e := range errs {
		field := e.Field()
		message := ""
		tag := e.Tag()

		switch e.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", field)
		case "email":
			message = fmt.Sprintf("%s is not a valid email", field)
		case "min":
			message = fmt.Sprintf("%s should be at least %s characters long", field, e.Param())
		}

		errors = append(errors, ErrorResponse{
			Tag:     tag,
			Field:   field,
			Message: message,
		})
	}

	return errors
}

var validate *validator.Validate

func ValidateFromStruct(data interface{}) interface{} {
	validate = validator.New()

	err := validate.Struct(data)

	if err != nil {
		errors := formatValidationErrors(err.(validator.ValidationErrors))

		return errors
	}

	return nil
}
