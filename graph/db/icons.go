package db

import (
	"gorm.io/gorm"
	"net/url"
	"time"
)

type Icons struct {
	ID        uint `gorm:"primaryKey;not null;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IconUrl   url.URL        `gorm:"not null"`
}
