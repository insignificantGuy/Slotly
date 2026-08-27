package models

import "gorm.io/gorm"

type CompanyRoleMapping struct {
	gorm.Model
	CompanyID string `gorm:"type:char(36);not null;uniqueIndex:idx_company_role"`
	RoleID    uint   `gorm:"not null;uniqueIndex:idx_company_role"`
}
