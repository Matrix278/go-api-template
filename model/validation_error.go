package model

import (
	"fmt"
	"go-api-template/pkg/json"
	"regexp"
	"strings"

	goccy "github.com/goccy/go-json"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Errors []ValidationErrorDetails `json:"errors"`
}

type ValidationErrorDetails struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func (validErr ValidationError) Error() string {
	bytes, _ := json.Encode(validErr)

	return string(bytes)
}

func ParseError(err error) ValidationError {
	var details []ValidationErrorDetails

	switch typedError := any(err).(type) {
	case validator.ValidationErrors:
		details = appendValidationErrors(typedError, details)
	case *goccy.UnmarshalTypeError:
		details = append(details, ValidationErrorDetails{
			Field:   toSnakeCase(typedError.Field),
			Message: parseMarshallingError(*typedError),
		})
	default:
		details = append(details, ValidationErrorDetails{
			Field:   "general",
			Message: err.Error(),
		})
	}

	return ValidationError{
		Errors: details,
	}
}

func appendValidationErrors(typedError validator.ValidationErrors, details []ValidationErrorDetails) []ValidationErrorDetails {
	for _, validationErr := range typedError {
		details = append(details, ValidationErrorDetails{
			Field:   toSnakeCase(validationErr.Field()),
			Message: parseFieldError(validationErr),
		})
	}

	return details
}

func parseFieldError(fieldError validator.FieldError) string {
	fieldPrefix := "the field " + toSnakeCase(fieldError.Field())

	switch fieldError.Tag() {
	case "required":
		return fieldPrefix + " is required"
	case "min":
		return fmt.Sprintf("%s should have at least %s elements", fieldPrefix, fieldError.Param())
	case "max":
		return fmt.Sprintf("%s should have at most %s elements", fieldPrefix, fieldError.Param())
	case "len":
		return fmt.Sprintf("%s should have a length of %s", fieldPrefix, fieldError.Param())
	case "phoneNumber":
		return fieldPrefix + " should be a valid phone number"
	case "latLong":
		return fieldPrefix + " should be a valid latitude and longitude"

	default:
		return fmt.Errorf("%w", fieldError).Error()
	}
}

func parseMarshallingError(unmarshalTypeError goccy.UnmarshalTypeError) string {
	return fmt.Sprintf("the field %s must be a %s", toSnakeCase(unmarshalTypeError.Field), unmarshalTypeError.Type.String())
}

func toSnakeCase(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(str, "${1}_${2}")

	return strings.ToLower(snake)
}
