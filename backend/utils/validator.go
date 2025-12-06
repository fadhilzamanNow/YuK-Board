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
			errors[field] = field + " wajib diisi"
		case "email":
			errors[field] = field + " harus berupa email yang valid"
		case "min":
			errors[field] = field + " minimal " + e.Param() + " karakter"
		case "max":
			errors[field] = field + " maksimal " + e.Param() + " karakter"
		case "oneof":
			errors[field] = field + " tidak valid"
		default:
			errors[field] = field + " tidak valid"
		}
	}

	return fiber.Map{
		"message": "Validasi gagal",
		"errors":  errors,
	}
}
