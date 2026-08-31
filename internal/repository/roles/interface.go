package roles

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
)

type RolesRepository interface {
	CreateRole(ctx context.Context, role *models.Role) error
	GetRole(ctx context.Context, id string) (*models.Role, error)
	UpdateRole(ctx context.Context, role *models.Role) error
	DeleteRole(ctx context.Context, id string) error
}
