package validator

import (
	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			field := fieldErr.Field()
			tag := fieldErr.Tag()

			switch tag {
			case "required":
				errors[field] = field + " is required"
			case "email":
				errors[field] = "Invalid email format"
			case "min":
				errors[field] = field + " is too short"
			case "max":
				errors[field] = field + " is too long"
			default:
				errors[field] = "Invalid value"
			}
		}
	} else {
		errors["error"] = err.Error()
	}

	return errors
}
