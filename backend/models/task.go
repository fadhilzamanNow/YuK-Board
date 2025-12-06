package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	TaskTodo       TaskStatus = "todo"
	TaskInProgress TaskStatus = "in_progress"
	TaskDone       TaskStatus = "done"
)

type Task struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	TodoListID  uuid.UUID  `gorm:"type:uuid;not null" json:"todo_list_id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `gorm:"type:task_status;default:'todo'" json:"status"`
	CreatedBy   uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	AssignedTo  *uuid.UUID `gorm:"type:uuid" json:"assigned_to,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	TodoList     TodoList `gorm:"foreignKey:TodoListID" json:"todo_list,omitempty"`
	Creator      User     `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	AssignedUser *User    `gorm:"foreignKey:AssignedTo" json:"assigned_user,omitempty"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}
