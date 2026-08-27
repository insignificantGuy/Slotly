package models

import "gorm.io/gorm"

type UserRoleMapping struct {
	gorm.Model
	UserID string `gorm:"type:char(36);not null;uniqueIndex:idx_user_role"`
	RoleID uint   `gorm:"not null;uniqueIndex:idx_user_role"`
}
