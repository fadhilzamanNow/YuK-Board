package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberRole string

const (
	RoleOwner  MemberRole = "owner"
	RoleMember MemberRole = "member"
)

type ListMember struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	TodoListID uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:uq_list_members" json:"todo_list_id"`
	UserID     uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:uq_list_members" json:"user_id"`
	Role       MemberRole `gorm:"type:member_role;not null" json:"role"`
	JoinedAt   time.Time  `gorm:"default:now()" json:"joined_at"`

	User     User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	TodoList TodoList `gorm:"foreignKey:TodoListID" json:"todo_list,omitempty"`
}

func (l *ListMember) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return nil
}
