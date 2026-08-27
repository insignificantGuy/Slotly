package models

import "time"

type Company struct {
	CompanyID string    `gorm:"type:char(36);default:(UUID());primaryKey"`
	Password  string    `gorm:"size:255"`
	Email     string    `gorm:"size:255;unique"`
	CreatedAt time.Time `gorm:"type:datetime(3);autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:datetime(3);autoUpdateTime"`
}
