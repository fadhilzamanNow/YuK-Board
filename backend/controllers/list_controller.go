package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"yukboard/config"
	"yukboard/models"
	"yukboard/utils"
)

type CreateListInput struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description"`
}

type UpdateListInput struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description"`
}

func CreateList(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var input CreateListInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input", "errors": nil})
	}

	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(400).JSON(utils.ValidationError(err))
	}

	list := models.TodoList{
		Title:       input.Title,
		Description: input.Description,
		OwnerID:     userID,
		InviteCode:  uuid.New().String()[:8],
	}

	if err := config.DB.Create(&list).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to create list", "errors": nil})
	}

	// Add owner to list_members
	member := models.ListMember{
		TodoListID: list.ID,
		UserID:     userID,
		Role:       models.RoleOwner,
	}
	config.DB.Create(&member)

	return c.Status(201).JSON(fiber.Map{"message": "List created", "list": list})
}

func GetLists(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var lists []models.TodoList
	config.DB.
		Joins("JOIN list_members ON list_members.todo_list_id = todo_lists.id").
		Where("list_members.user_id = ?", userID).
		Preload("Owner").
		Find(&lists)

	return c.JSON(fiber.Map{"message": "Success", "lists": lists})
}

func GetList(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid list ID", "errors": nil})
	}

	// Check membership
	var member models.ListMember
	if err := config.DB.Where("todo_list_id = ? AND user_id = ?", listID, userID).First(&member).Error; err != nil {
		return c.Status(403).JSON(fiber.Map{"message": "Access denied", "errors": nil})
	}

	var list models.TodoList
	config.DB.Preload("Owner").Preload("Members.User").First(&list, "id = ?", listID)

	return c.JSON(fiber.Map{"message": "Success", "list": list})
}

func UpdateList(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid list ID", "errors": nil})
	}

	// Check owner
	var list models.TodoList
	if err := config.DB.First(&list, "id = ? AND owner_id = ?", listID, userID).Error; err != nil {
		return c.Status(403).JSON(fiber.Map{"message": "Access denied", "errors": nil})
	}

	var input UpdateListInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input", "errors": nil})
	}

	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(400).JSON(utils.ValidationError(err))
	}

	list.Title = input.Title
	list.Description = input.Description
	config.DB.Save(&list)

	return c.JSON(fiber.Map{"message": "List updated", "list": list})
}

func DeleteList(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid list ID", "errors": nil})
	}

	// First check if list exists and user is owner
	var list models.TodoList
	if err := config.DB.First(&list, "id = ?", listID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "List not found", "errors": nil})
	}

	if list.OwnerID != userID {
		return c.Status(403).JSON(fiber.Map{"message": "Access denied", "errors": nil})
	}

	// Delete related records first
	config.DB.Where("todo_list_id = ?", list.ID).Delete(&models.ListMember{})
	config.DB.Where("todo_list_id = ?", list.ID).Delete(&models.ListInvitation{})
	config.DB.Where("todo_list_id = ?", list.ID).Delete(&models.Task{})
	config.DB.Unscoped().Delete(&list)

	return c.JSON(fiber.Map{"message": "List deleted"})
}
