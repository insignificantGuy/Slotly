package models

import "gorm.io/gorm"

const (
	BookingStatusConfirmed = "confirmed"
	BookingStatusCancelled = "cancelled"
)

type Booking struct {
	gorm.Model
	UserID    string `gorm:"type:char(36);not null;index"`
	SlotID    uint   `gorm:"not null;uniqueIndex"`
	ListingID string `gorm:"type:char(36);not null;index"`
	Status    string `gorm:"size:32;not null;default:confirmed"`
}
