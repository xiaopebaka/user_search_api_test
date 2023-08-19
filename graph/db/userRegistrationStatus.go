package db

import (
	"gorm.io/gorm"
	"time"
)

type UserRegistrationStatus struct {
	ID              uint      `gorm:"primaryKey;not null;autoIncrement"`
	AccountID       string    `gorm:"not null"`
	PreRegisteredAt time.Time `gorm:"not null"`
	RegisteredAt    *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
