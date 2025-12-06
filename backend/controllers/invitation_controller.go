package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"yukboard/config"
	"yukboard/models"
	"yukboard/utils"
)

type InviteInput struct {
	Email string `json:"email" validate:"required,email"`
}

func InviteUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID daftar tidak valid", "errors": nil})
	}

	var list models.TodoList
	if err := config.DB.First(&list, "id = ? AND owner_id = ?", listID, userID).Error; err != nil {
		return c.Status(403).JSON(fiber.Map{"message": "Akses ditolak", "errors": nil})
	}

	var input InviteInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Input tidak valid", "errors": nil})
	}

	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(400).JSON(utils.ValidationError(err))
	}

	var invitee models.User
	if err := config.DB.First(&invitee, "email = ?", input.Email).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Pengguna tidak ditemukan", "errors": nil})
	}

	if invitee.ID == userID {
		return c.Status(400).JSON(fiber.Map{"message": "Tidak dapat mengundang diri sendiri", "errors": nil})
	}

	var existingMember models.ListMember
	if err := config.DB.First(&existingMember, "todo_list_id = ? AND user_id = ?", listID, invitee.ID).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"message": "Pengguna sudah menjadi anggota", "errors": nil})
	}

	var existingInvite models.ListInvitation
	if err := config.DB.First(&existingInvite, "todo_list_id = ? AND invitee_id = ? AND status = ?", listID, invitee.ID, models.StatusPending).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"message": "Undangan sudah dikirim", "errors": nil})
	}

	invitation := models.ListInvitation{
		TodoListID: listID,
		InviterID:  userID,
		InviteeID:  invitee.ID,
		Status:     models.StatusPending,
		Token:      uuid.New().String(),
	}
	config.DB.Create(&invitation)

	return c.Status(201).JSON(fiber.Map{"message": "Undangan berhasil dikirim", "invitation": invitation})
}

func GetMyInvitations(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var invitations []models.ListInvitation
	config.DB.
		Where("invitee_id = ? AND status = ?", userID, models.StatusPending).
		Preload("TodoList").
		Preload("Inviter").
		Find(&invitations)

	return c.JSON(fiber.Map{"message": "Berhasil", "invitations": invitations})
}

func AcceptInvitation(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	inviteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID undangan tidak valid", "errors": nil})
	}

	var invitation models.ListInvitation
	if err := config.DB.First(&invitation, "id = ? AND invitee_id = ? AND status = ?", inviteID, userID, models.StatusPending).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Undangan tidak ditemukan", "errors": nil})
	}

	now := time.Now()
	invitation.Status = models.StatusAccepted
	invitation.RespondedAt = &now
	config.DB.Save(&invitation)

	member := models.ListMember{
		TodoListID: invitation.TodoListID,
		UserID:     userID,
		Role:       models.RoleMember,
	}
	config.DB.Create(&member)

	return c.JSON(fiber.Map{"message": "Undangan diterima"})
}

func DeclineInvitation(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	inviteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID undangan tidak valid", "errors": nil})
	}

	var invitation models.ListInvitation
	if err := config.DB.First(&invitation, "id = ? AND invitee_id = ? AND status = ?", inviteID, userID, models.StatusPending).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Undangan tidak ditemukan", "errors": nil})
	}

	now := time.Now()
	invitation.Status = models.StatusDeclined
	invitation.RespondedAt = &now
	config.DB.Save(&invitation)

	return c.JSON(fiber.Map{"message": "Undangan ditolak"})
}

func CancelInvitation(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	inviteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "ID undangan tidak valid", "errors": nil})
	}

	var invitation models.ListInvitation
	if err := config.DB.First(&invitation, "id = ? AND inviter_id = ? AND status = ?", inviteID, userID, models.StatusPending).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Undangan tidak ditemukan", "errors": nil})
	}

	now := time.Now()
	invitation.Status = models.StatusCancelled
	invitation.RespondedAt = &now
	config.DB.Save(&invitation)

	return c.JSON(fiber.Map{"message": "Undangan dibatalkan"})
}
