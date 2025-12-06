package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"yukboard/config"
	"yukboard/models"
	"yukboard/utils"
)

type CreateTaskInput struct {
	Title       string     `json:"title" validate:"required,min=3"`
	Description string     `json:"description"`
	AssignedTo  *uuid.UUID `json:"assigned_to"`
}

type UpdateTaskInput struct {
	Title       string     `json:"title" validate:"required,min=3"`
	Description string     `json:"description"`
	Status      string     `json:"status" validate:"omitempty,oneof=todo in_progress done"`
	AssignedTo  *uuid.UUID `json:"assigned_to"`
}

// Check if user is member of list
func isMember(listID, userID uuid.UUID) bool {
	var member models.ListMember
	return config.DB.First(&member, "todo_list_id = ? AND user_id = ?", listID, userID).Error == nil
}

// Check if user is owner of list
func isOwner(listID, userID uuid.UUID) bool {
	var list models.TodoList
	return config.DB.First(&list, "id = ? AND owner_id = ?", listID, userID).Error == nil
}

func CreateTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid list ID", "errors": nil})
	}

	if !isMember(listID, userID) {
		return c.Status(403).JSON(fiber.Map{"message": "Access denied", "errors": nil})
	}

	var input CreateTaskInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input", "errors": nil})
	}

	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(400).JSON(utils.ValidationError(err))
	}

	// Validate assigned_to is a member
	if input.AssignedTo != nil && !isMember(listID, *input.AssignedTo) {
		return c.Status(400).JSON(fiber.Map{"message": "Assigned user is not a member", "errors": nil})
	}

	task := models.Task{
		TodoListID:  listID,
		Title:       input.Title,
		Description: input.Description,
		Status:      models.TaskTodo,
		CreatedBy:   userID,
		AssignedTo:  input.AssignedTo,
	}
	config.DB.Create(&task)

	return c.Status(201).JSON(fiber.Map{"message": "Task created", "task": task})
}

func GetTasks(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid list ID", "errors": nil})
	}

	if !isMember(listID, userID) {
		return c.Status(403).JSON(fiber.Map{"message": "Access denied", "errors": nil})
	}

	var tasks []models.Task
	config.DB.Where("todo_list_id = ?", listID).Preload("Creator").Preload("AssignedUser").Find(&tasks)

	return c.JSON(fiber.Map{"message": "Success", "tasks": tasks})
}

func GetTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid list ID", "errors": nil})
	}
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid task ID", "errors": nil})
	}

	if !isMember(listID, userID) {
		return c.Status(403).JSON(fiber.Map{"message": "Access denied", "errors": nil})
	}

	var task models.Task
	if err := config.DB.Preload("Creator").Preload("AssignedUser").First(&task, "id = ? AND todo_list_id = ?", taskID, listID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found", "errors": nil})
	}

	return c.JSON(fiber.Map{"message": "Success", "task": task})
}

func UpdateTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid list ID", "errors": nil})
	}
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid task ID", "errors": nil})
	}

	var task models.Task
	if err := config.DB.First(&task, "id = ? AND todo_list_id = ?", taskID, listID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found", "errors": nil})
	}

	// Owner can edit any task, member can only edit their own
	if !isOwner(listID, userID) && task.CreatedBy != userID {
		return c.Status(403).JSON(fiber.Map{"message": "Access denied", "errors": nil})
	}

	var input UpdateTaskInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input", "errors": nil})
	}

	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(400).JSON(utils.ValidationError(err))
	}

	if input.AssignedTo != nil && !isMember(listID, *input.AssignedTo) {
		return c.Status(400).JSON(fiber.Map{"message": "Assigned user is not a member", "errors": nil})
	}

	task.Title = input.Title
	task.Description = input.Description
	if input.Status != "" {
		task.Status = models.TaskStatus(input.Status)
	}
	task.AssignedTo = input.AssignedTo
	config.DB.Save(&task)

	return c.JSON(fiber.Map{"message": "Task updated", "task": task})
}

func DeleteTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid list ID", "errors": nil})
	}
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid task ID", "errors": nil})
	}

	var task models.Task
	if err := config.DB.First(&task, "id = ? AND todo_list_id = ?", taskID, listID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found", "errors": nil})
	}

	// Owner can delete any task, member can delete their own
	if !isOwner(listID, userID) && task.CreatedBy != userID {
		return c.Status(403).JSON(fiber.Map{"message": "Access denied", "errors": nil})
	}

	config.DB.Delete(&task)

	return c.JSON(fiber.Map{"message": "Task deleted"})
}

func UpdateTaskStatus(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid list ID", "errors": nil})
	}
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid task ID", "errors": nil})
	}

	if !isMember(listID, userID) {
		return c.Status(403).JSON(fiber.Map{"message": "Access denied", "errors": nil})
	}

	var task models.Task
	if err := config.DB.First(&task, "id = ? AND todo_list_id = ?", taskID, listID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Task not found", "errors": nil})
	}

	// Cycle status: todo -> in_progress -> done -> todo
	switch task.Status {
	case models.TaskTodo:
		task.Status = models.TaskInProgress
	case models.TaskInProgress:
		task.Status = models.TaskDone
	case models.TaskDone:
		task.Status = models.TaskTodo
	}
	config.DB.Save(&task)

	return c.JSON(fiber.Map{"message": "Status updated", "task": task})
}
