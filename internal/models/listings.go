package models

import (
	"time"

	"gorm.io/gorm"
)

type Listing struct {
	ListingID   string         `gorm:"column:listing_id;type:char(36);default:(UUID());primaryKey"`
	Type        string         `gorm:"not null"`
	Title       string         `gorm:"not null"`
	Image       string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	CompanyID   string         `gorm:"type:char(36);not null"`
	Price       float64        `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
