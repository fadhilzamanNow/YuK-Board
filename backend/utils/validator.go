package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

func ValidationError(err error) fiber.Map {
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		field := strings.ToLower(e.Field())
		switch e.Tag() {
		case "required":
			errors[field] = field + " is required"
		case "email":
			errors[field] = field + " must be a valid email"
		case "min":
			errors[field] = field + " must be at least " + e.Param() + " characters"
		case "max":
			errors[field] = field + " must be at most " + e.Param() + " characters"
		default:
			errors[field] = field + " is invalid"
		}
	}

	return fiber.Map{
		"message": "Validation failed",
		"errors":  errors,
	}
}
