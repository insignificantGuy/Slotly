package models

import (
	"time"
)

type Listing struct {
	ListingID   string    `gorm:"primaryKey"`
	Type        string    `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Image       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	CompanyID   Company   `foreignKey:"CompanyID" gorm:"not null"`
	Price       float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	DeletedAt   time.Time `gorm:"not null" softDelete:"true"`
}
