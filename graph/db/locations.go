package db

import (
	"gorm.io/gorm"
	"time"
)

type Locations struct {
	ID           uint `gorm:"primaryKey;not null;autoIncrement"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	LocationName string         `gorm:"not null"`
}
