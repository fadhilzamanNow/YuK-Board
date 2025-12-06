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

func isMember(listID, userID uuid.UUID) bool {
	var member models.ListMember
	return config.DB.First(&member, "todo_list_id = ? AND user_id = ?", listID, userID).Error == nil
}

func isOwner(listID, userID uuid.UUID) bool {
	var list models.TodoList
	return config.DB.First(&list, "id = ? AND owner_id = ?", listID, userID).Error == nil
}

func CreateTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID daftar tidak valid", "errors": nil})
	}

	if !isMember(listID, userID) {
		return c.Status(403).JSON(fiber.Map{"message": "Akses ditolak", "errors": nil})
	}

	var input CreateTaskInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Input tidak valid", "errors": nil})
	}

	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(400).JSON(utils.ValidationError(err))
	}

	if input.AssignedTo != nil && !isMember(listID, *input.AssignedTo) {
		return c.Status(400).JSON(fiber.Map{"message": "Pengguna yang ditugaskan bukan anggota", "errors": nil})
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

	return c.Status(201).JSON(fiber.Map{"message": "Tugas berhasil dibuat", "task": task})
}

func GetTasks(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID daftar tidak valid", "errors": nil})
	}

	if !isMember(listID, userID) {
		return c.Status(403).JSON(fiber.Map{"message": "Akses ditolak", "errors": nil})
	}

	var tasks []models.Task
	config.DB.Where("todo_list_id = ?", listID).Preload("Creator").Preload("AssignedUser").Find(&tasks)

	return c.JSON(fiber.Map{"message": "Berhasil", "tasks": tasks})
}

func GetTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID daftar tidak valid", "errors": nil})
	}
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID tugas tidak valid", "errors": nil})
	}

	if !isMember(listID, userID) {
		return c.Status(403).JSON(fiber.Map{"message": "Akses ditolak", "errors": nil})
	}

	var task models.Task
	if err := config.DB.Preload("Creator").Preload("AssignedUser").First(&task, "id = ? AND todo_list_id = ?", taskID, listID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Tugas tidak ditemukan", "errors": nil})
	}

	return c.JSON(fiber.Map{"message": "Berhasil", "task": task})
}

func UpdateTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID daftar tidak valid", "errors": nil})
	}
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID tugas tidak valid", "errors": nil})
	}

	var task models.Task
	if err := config.DB.First(&task, "id = ? AND todo_list_id = ?", taskID, listID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Tugas tidak ditemukan", "errors": nil})
	}

	if !isOwner(listID, userID) && task.CreatedBy != userID {
		return c.Status(403).JSON(fiber.Map{"message": "Akses ditolak", "errors": nil})
	}

	var input UpdateTaskInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Input tidak valid", "errors": nil})
	}

	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(400).JSON(utils.ValidationError(err))
	}

	if input.AssignedTo != nil && !isMember(listID, *input.AssignedTo) {
		return c.Status(400).JSON(fiber.Map{"message": "Pengguna yang ditugaskan bukan anggota", "errors": nil})
	}

	task.Title = input.Title
	task.Description = input.Description
	if input.Status != "" {
		task.Status = models.TaskStatus(input.Status)
	}
	task.AssignedTo = input.AssignedTo
	config.DB.Save(&task)

	return c.JSON(fiber.Map{"message": "Tugas berhasil diperbarui", "task": task})
}

func DeleteTask(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID daftar tidak valid", "errors": nil})
	}
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID tugas tidak valid", "errors": nil})
	}

	var task models.Task
	if err := config.DB.First(&task, "id = ? AND todo_list_id = ?", taskID, listID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Tugas tidak ditemukan", "errors": nil})
	}

	if !isOwner(listID, userID) && task.CreatedBy != userID {
		return c.Status(403).JSON(fiber.Map{"message": "Akses ditolak", "errors": nil})
	}

	config.DB.Delete(&task)

	return c.JSON(fiber.Map{"message": "Tugas berhasil dihapus"})
}

func UpdateTaskStatus(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("listId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID daftar tidak valid", "errors": nil})
	}
	taskID, err := uuid.Parse(c.Params("taskId"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID tugas tidak valid", "errors": nil})
	}

	if !isMember(listID, userID) {
		return c.Status(403).JSON(fiber.Map{"message": "Akses ditolak", "errors": nil})
	}

	var task models.Task
	if err := config.DB.First(&task, "id = ? AND todo_list_id = ?", taskID, listID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Tugas tidak ditemukan", "errors": nil})
	}

	switch task.Status {
	case models.TaskTodo:
		task.Status = models.TaskInProgress
	case models.TaskInProgress:
		task.Status = models.TaskDone
	case models.TaskDone:
		task.Status = models.TaskTodo
	}
	config.DB.Save(&task)

	return c.JSON(fiber.Map{"message": "Status berhasil diperbarui", "task": task})
}
