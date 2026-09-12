package refreshtoken

import (
	"context"
	"errors"

	"github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"gorm.io/gorm"
)

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *refreshTokenRepository {
	return &refreshTokenRepository{db: db}
}

func (r *refreshTokenRepository) Create(ctx context.Context, token *models.RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *refreshTokenRepository) GetByHash(ctx context.Context, hash string) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := r.db.WithContext(ctx).Where("token_hash = ?", hash).First(&token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &token, nil
}

func (r *refreshTokenRepository) Revoke(ctx context.Context, hash string) error {
	return r.db.WithContext(ctx).Model(&models.RefreshToken{}).Where("token_hash = ?", hash).Update("revoked", true).Error
}

func (r *refreshTokenRepository) RevokeAllForUser(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).Model(&models.RefreshToken{}).Where("user_id = ?", userID).Update("revoked", true).Error
}
