package db

import (
	"gorm.io/gorm"
	"time"
)

type Locations struct {
	ID           string `gorm:"primaryKey;not null;type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	LocationName string         `gorm:"not null"`
}
