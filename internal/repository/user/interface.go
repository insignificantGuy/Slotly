package user

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	UpdateUserFields(ctx context.Context, id string, fields map[string]any) error
	DeleteUser(ctx context.Context, id string) error
}
