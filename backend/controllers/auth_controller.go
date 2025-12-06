package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"yukboard/config"
	"yukboard/models"
	"yukboard/utils"
)

type RegisterInput struct {
	Name     string `json:"name" validate:"required,min=5"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Register(c *fiber.Ctx) error {
	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input", "errors": nil})
	}

	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(400).JSON(utils.ValidationError(err))
	}

	hash, _ := utils.HashPassword(input.Password)
	user := models.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hash,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
			"errors":  fiber.Map{"Email": "Email already exists"},
		})
	}

	token, _ := utils.GenerateToken(user.ID)
	return c.Status(201).JSON(fiber.Map{"message": "Registration successful", "token": token, "user": user})
}

func Login(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input", "errors": nil})
	}

	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(400).JSON(utils.ValidationError(err))
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Your email or password is wrong", "errors": nil})
	}

	if !utils.CheckPassword(input.Password, user.PasswordHash) {
		return c.Status(401).JSON(fiber.Map{"message": "Your email or password is wrong", "errors": nil})
	}

	token, _ := utils.GenerateToken(user.ID)
	return c.JSON(fiber.Map{"message": "Login successful", "token": token, "user": user})
}

func Me(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var user models.User
	if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found", "errors": nil})
	}

	return c.JSON(fiber.Map{"message": "Success", "user": user})
}
