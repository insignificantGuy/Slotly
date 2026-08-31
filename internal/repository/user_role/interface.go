package userrole

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
)

type UserRoleRepository interface {
	CreateUserRole(ctx context.Context, userRole *models.UserRoleMapping) error
	GetUserRole(ctx context.Context, id string) (*models.UserRoleMapping, error)
	UpdateUserRole(ctx context.Context, userRole *models.UserRoleMapping) error
	DeleteUserRole(ctx context.Context, id string) error
}
