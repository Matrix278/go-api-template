package model

import (
	"fmt"
	"regexp"
	"strings"
	"go-api-template/pkg/json"

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
	fieldPrefix := fmt.Sprintf("the field %s", toSnakeCase(fieldError.Field()))
	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fieldPrefix)
	case "policy":
		return fmt.Sprintf("%s should be a valid policy", fieldPrefix)
	case "policies":
		return fmt.Sprintf("%s should be a valid policy", fieldPrefix)
	case "subject":
		return fmt.Sprintf("%s should be an existing user", fieldPrefix)
	case "object":
		return fmt.Sprintf("%s should be an existing endpoint", fieldPrefix)
	case "action":
		return fmt.Sprintf("%s should be a valid HTTP method", fieldPrefix)
	case "currencyCodeISO4217":
		return fmt.Sprintf("%s should be a valid ISO 4217 currency code", fieldPrefix)
	case "amadeusSource":
		return fmt.Sprintf("%s should be a valid Amadeus source", fieldPrefix)
	case "min":
		return fmt.Sprintf("%s should have at least %s elements", fieldPrefix, fieldError.Param())
	case "max":
		return fmt.Sprintf("%s should have at most %s elements", fieldPrefix, fieldError.Param())
	case "dive":
		return fmt.Sprintf("%s should be a valid %s", fieldPrefix, fieldError.Param())
	case "iatacode":
		return fmt.Sprintf("%s should be a valid IATA code", fieldPrefix)
	case "datetime":
		return fmt.Sprintf("%s should be a valid date", fieldPrefix)
	case "travelClass":
		return fmt.Sprintf("%s should be a valid travel class", fieldPrefix)
	case "len":
		return fmt.Sprintf("%s should have a length of %s", fieldPrefix, fieldError.Param())
	case "departureDate":
		return fmt.Sprintf("%s should be a valid departure date", fieldPrefix)
	case "duration":
		return fmt.Sprintf("%s should be a valid duration", fieldPrefix)
	case "viewBy":
		return fmt.Sprintf("%s should be a valid view by", fieldPrefix)
	case "tripAdvisorLanguage":
		return fmt.Sprintf("%s should be a valid TripAdvisor language", fieldPrefix)
	case "tripAdvisorSource":
		return fmt.Sprintf("%s should be a valid TripAdvisor source", fieldPrefix)
	case "tripAdvisorCategory":
		return fmt.Sprintf("%s should be a valid TripAdvisor category", fieldPrefix)
	case "phoneNumber":
		return fmt.Sprintf("%s should be a valid phone number", fieldPrefix)
	case "latLong":
		return fmt.Sprintf("%s should be a valid latitude and longitude", fieldPrefix)
	case "tripAdvisorRadiusUnit":
		return fmt.Sprintf("%s should be a valid TripAdvisor radius unit", fieldPrefix)

	default:
		return fmt.Errorf("%v", fieldError).Error()
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
