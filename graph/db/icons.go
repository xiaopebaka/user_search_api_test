package db

import (
	"gorm.io/gorm"
	"net/url"
	"time"
)

type Icons struct {
	ID        string `gorm:"primaryKey;not null;type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IconUrl   url.URL        `gorm:"not null"`
}
