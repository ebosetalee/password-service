package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "min":
		return fmt.Sprintf("must be at least %s characters long", fe.Param())
	}
	return fe.Error() // default error

}

func ValidateError(err error) string {
	ve := err.(validator.ValidationErrors)
	if errors.As(err, &ve) {
		var errorDetails string
		for _, fe := range ve {
			errorDetails += fmt.Sprintf(" %s %s", fe.Field(), msgForTag(fe))
		}
		return "Invalid input:" + errorDetails
	}
	return ""
}
