package models

import "gorm.io/gorm"

type CompanyRoleMapping struct {
	gorm.Model
	UserID    string `gorm:"type:char(36);not null;uniqueIndex:idx_user_company"`
	CompanyID string `gorm:"type:char(36);not null;uniqueIndex:idx_user_company"`
	RoleID    uint   `gorm:"not null"`
}
