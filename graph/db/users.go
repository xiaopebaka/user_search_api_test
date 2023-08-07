package db

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID         string `gorm:"primaryKey;not null;type:varchar(255)"`
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
