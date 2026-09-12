package refreshtoken

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, token *models.RefreshToken) error
	GetByHash(ctx context.Context, hash string) (*models.RefreshToken, error)
	Revoke(ctx context.Context, hash string) error
	RevokeAllForUser(ctx context.Context, userID string) error
}
