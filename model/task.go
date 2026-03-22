package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	StatusPending    Status = "pending"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
	StatusExpired    Status = "expired"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Status      Status    `gorm:"type:varchar(255);default:'pending'" json:"status"`
	CreatedAt   int64     `json:"createdAt"`
	ExpiresAt   int64     `gorm:"not null" json:"expiresAt"`
}

type CreateTaskRequest struct {
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	ExpiresAt   int64  `gorm:"not null" json:"expiresAt"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	ExpiresAt   int64  `json:"expiresAt"`
}

func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	task.ID = uuid.New()
	task.CreatedAt = time.Now().Unix()

	if task.ExpiresAt <= task.CreatedAt {
		task.Status = StatusExpired
	}

	return
}
