package models

import (
	"time"
)

type User struct {
	UserID      string    `gorm:"type:char(36);default:(UUID());primaryKey"`
	FullName    string    `gorm:"size:255"`
	Email       string    `gorm:"size:255;unique"`
	Password    string    `gorm:"size:255"`
	PhoneNumber string    `gorm:"size:20;unique"`
	CreatedAt   time.Time `gorm:"type:datetime(3);autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:datetime(3);autoUpdateTime"`
}
