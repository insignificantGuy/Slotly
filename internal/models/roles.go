package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name    string `gorm:"size:255;unique"`
	Actions string `gorm:"size:255"`
}
