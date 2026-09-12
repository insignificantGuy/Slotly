package models

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    string    `gorm:"type:char(36);not null;index"`
	TokenHash string    `gorm:"size:64;not null;uniqueIndex"`
	ExpiresAt time.Time `gorm:"type:datetime(3);not null"`
	Revoked   bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
