package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	TaskId          uuid.UUID `gorm:"primaryKey"`
	TaskName        string    `gorm:"not null"`
	TaskDescription string    `gorm:"not null"`
	TaskStatus      string    `gorm:"not null"`
	TaskCreatedAt   time.Time `gorm:"not null"`
	TaskUpdatedAt   time.Time `gorm:"not null"`
}
