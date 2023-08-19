package db

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID         uint   `gorm:"primaryKey;not null;autoIncrements"`
	AccountID  string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Name       string
	Birth      *time.Time
	IconId     *string
	Icons      Icons `gorm:"foreignKey:IconId"`
	Profile    *string
	LocationId *string
	Locations  Locations `gorm:"foreignKey:LocationId"`
}
