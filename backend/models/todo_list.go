package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoList struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	OwnerID     uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
	InviteCode  string    `gorm:"not null;uniqueIndex" json:"invite_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Owner   User         `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
	Members []ListMember `gorm:"foreignKey:TodoListID" json:"members,omitempty"`
}

func (t *TodoList) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}
