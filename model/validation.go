package model

import (
	"go-api-template/pkg/logger"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var (
	phoneRegex   = regexp.MustCompile(`^[0-9- ]+$`)
	latLongRegex = regexp.MustCompile(`\s*,\s*`)
)

func InitValidation() error {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Custom validations
		validationRules := []struct {
			Type string
			Func validator.Func
		}{
			{"datetime", validateDatetime},
			{"phoneNumber", validatePhoneNumber},
			{"latLong", validateLatLong},
		}

		for _, validationRule := range validationRules {
			if err := registerValidation(validate, validationRule.Type, validationRule.Func); err != nil {
				return err
			}
		}
	}

	return nil
}

// private.
func registerValidation(validate *validator.Validate, validationType string, validationFunc validator.Func) error {
	if err := validate.RegisterValidation(validationType, validationFunc); err != nil {
		logger.Errorf("error registering %s validation: %v", validationType, err)

		return err
	}

	return nil
}

func validateDatetime(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())

	return err == nil
}

func validatePhoneNumber(fl validator.FieldLevel) bool {
	return phoneRegex.MatchString(fl.Field().String())
}

func validateLatLong(fl validator.FieldLevel) bool {
	latLongSlice := latLongRegex.Split(fl.Field().String(), -1)

	if len(latLongSlice) != 2 {
		return false
	}

	latitude, err := strconv.ParseFloat(latLongSlice[0], 64)
	if err != nil {
		return false
	}

	longitude, err := strconv.ParseFloat(latLongSlice[1], 64)
	if err != nil {
		return false
	}

	return latitude >= -90 && latitude <= 90 && longitude >= -180 && longitude <= 180
}
