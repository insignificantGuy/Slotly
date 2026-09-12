package models

import "time"

type Company struct {
	CompanyID   string    `gorm:"type:char(36);default:(UUID());primaryKey"`
	Name        string    `gorm:"size:255;not null"`
	Address     string    `gorm:"size:255;not null"`
	Phone       string    `gorm:"size:255;not null"`
	Email       string    `gorm:"size:255;not null"`
	Logo        string    `gorm:"size:255;not null"`
	Description string    `gorm:"size:255;not null"`
	IsActive    bool      `gorm:"default:true"`
	IsDeleted   bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"type:datetime(3);autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:datetime(3);autoUpdateTime"`
}
