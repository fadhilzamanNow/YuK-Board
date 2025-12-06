package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvitationStatus string

const (
	StatusPending   InvitationStatus = "pending"
	StatusAccepted  InvitationStatus = "accepted"
	StatusDeclined  InvitationStatus = "declined"
	StatusCancelled InvitationStatus = "cancelled"
)

type ListInvitation struct {
	ID          uuid.UUID        `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	TodoListID  uuid.UUID        `gorm:"type:uuid;not null" json:"todo_list_id"`
	InviterID   uuid.UUID        `gorm:"type:uuid;not null" json:"inviter_id"`
	InviteeID   uuid.UUID        `gorm:"type:uuid;not null" json:"invitee_id"`
	Status      InvitationStatus `gorm:"type:invitation_status;default:'pending'" json:"status"`
	Token       string           `gorm:"not null;uniqueIndex" json:"token"`
	CreatedAt   time.Time        `json:"created_at"`
	RespondedAt *time.Time       `json:"responded_at,omitempty"`

	TodoList TodoList `gorm:"foreignKey:TodoListID" json:"todo_list,omitempty"`
	Inviter  User     `gorm:"foreignKey:InviterID" json:"inviter,omitempty"`
	Invitee  User     `gorm:"foreignKey:InviteeID" json:"invitee,omitempty"`
}

func (i *ListInvitation) BeforeCreate(tx *gorm.DB) error {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return nil
}
