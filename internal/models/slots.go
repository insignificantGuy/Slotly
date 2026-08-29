package models

import (
	"time"

	"gorm.io/gorm"
)

type Slot struct {
	gorm.Model
	ListingID Listing   `foreignKey:"ListingID" gorm:"not null"`
	Date      time.Time `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
	Duration  int       `gorm:"not null"`
	IsBooked  bool      `gorm:"default:false"`
}
