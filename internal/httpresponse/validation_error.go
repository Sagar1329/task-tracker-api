package httpresponse

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)
func splitCamelCase(s string) string {

	var result []rune

	for i, r := range s {

		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, ' ')
		}

		result = append(result, r)
	}

	return string(result)
}
func FormatValidationError(err error) string {

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	validationError := validationErrors[0]

	field := splitCamelCase(validationError.Field())

	switch validationError.Tag() {

	case "required":
		return fmt.Sprintf("%s is required.", field)

	case "email":
		return fmt.Sprintf("%s must be a valid email address.", field)

	case "url":
		return fmt.Sprintf("%s must be a valid URL.", field)

	case "min":
		return fmt.Sprintf(
			"%s must be at least %s characters.",
			field,
			validationError.Param(),
		)

	case "max":
		return fmt.Sprintf(
			"%s must not exceed %s characters.",
			field,
			validationError.Param(),
		)

	case "oneof":
		return fmt.Sprintf(
			"%s must be one of: %s.",
			field,
			strings.ReplaceAll(validationError.Param(), " ", ", "),
		)

	case "datetime":
		return fmt.Sprintf(
			"%s must be in YYYY-MM-DD format.",
			field,
		)

	default:
		return fmt.Sprintf("%s is invalid.", field)
	}
}