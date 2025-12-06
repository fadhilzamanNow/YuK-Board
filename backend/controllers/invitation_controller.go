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

// Owner invites user by email
func InviteUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	listID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid list ID", "errors": nil})
	}

	// Check if user is owner
	var list models.TodoList
	if err := config.DB.First(&list, "id = ? AND owner_id = ?", listID, userID).Error; err != nil {
		return c.Status(403).JSON(fiber.Map{"message": "Access denied", "errors": nil})
	}

	var input InviteInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid input", "errors": nil})
	}

	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(400).JSON(utils.ValidationError(err))
	}

	// Find invitee by email
	var invitee models.User
	if err := config.DB.First(&invitee, "email = ?", input.Email).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found", "errors": nil})
	}

	// Can't invite yourself
	if invitee.ID == userID {
		return c.Status(400).JSON(fiber.Map{"message": "Cannot invite yourself", "errors": nil})
	}

	// Check if already a member
	var existingMember models.ListMember
	if err := config.DB.First(&existingMember, "todo_list_id = ? AND user_id = ?", listID, invitee.ID).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"message": "User is already a member", "errors": nil})
	}

	// Check if pending invitation exists
	var existingInvite models.ListInvitation
	if err := config.DB.First(&existingInvite, "todo_list_id = ? AND invitee_id = ? AND status = ?", listID, invitee.ID, models.StatusPending).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invitation already sent", "errors": nil})
	}

	invitation := models.ListInvitation{
		TodoListID: listID,
		InviterID:  userID,
		InviteeID:  invitee.ID,
		Status:     models.StatusPending,
		Token:      uuid.New().String(),
	}
	config.DB.Create(&invitation)

	return c.Status(201).JSON(fiber.Map{"message": "Invitation sent", "invitation": invitation})
}

// Get pending invitations for current user
func GetMyInvitations(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var invitations []models.ListInvitation
	config.DB.
		Where("invitee_id = ? AND status = ?", userID, models.StatusPending).
		Preload("TodoList").
		Preload("Inviter").
		Find(&invitations)

	return c.JSON(fiber.Map{"message": "Success", "invitations": invitations})
}

// Accept invitation
func AcceptInvitation(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	inviteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid invitation ID", "errors": nil})
	}

	var invitation models.ListInvitation
	if err := config.DB.First(&invitation, "id = ? AND invitee_id = ? AND status = ?", inviteID, userID, models.StatusPending).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Invitation not found", "errors": nil})
	}

	// Update invitation status
	now := time.Now()
	invitation.Status = models.StatusAccepted
	invitation.RespondedAt = &now
	config.DB.Save(&invitation)

	// Add user to list members
	member := models.ListMember{
		TodoListID: invitation.TodoListID,
		UserID:     userID,
		Role:       models.RoleMember,
	}
	config.DB.Create(&member)

	return c.JSON(fiber.Map{"message": "Invitation accepted"})
}

// Decline invitation
func DeclineInvitation(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	inviteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid invitation ID", "errors": nil})
	}

	var invitation models.ListInvitation
	if err := config.DB.First(&invitation, "id = ? AND invitee_id = ? AND status = ?", inviteID, userID, models.StatusPending).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Invitation not found", "errors": nil})
	}

	now := time.Now()
	invitation.Status = models.StatusDeclined
	invitation.RespondedAt = &now
	config.DB.Save(&invitation)

	return c.JSON(fiber.Map{"message": "Invitation declined"})
}

// Owner cancels invitation
func CancelInvitation(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)
	inviteID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid invitation ID", "errors": nil})
	}

	var invitation models.ListInvitation
	if err := config.DB.First(&invitation, "id = ? AND inviter_id = ? AND status = ?", inviteID, userID, models.StatusPending).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Invitation not found", "errors": nil})
	}

	now := time.Now()
	invitation.Status = models.StatusCancelled
	invitation.RespondedAt = &now
	config.DB.Save(&invitation)

	return c.JSON(fiber.Map{"message": "Invitation cancelled"})
}
