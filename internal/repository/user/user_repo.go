package user

import (
	"context"
	"errors"

	"github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	*repository.BaseRepository[models.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		BaseRepository: repository.NewBaseRepository[models.User](db),
		db:             db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.BaseRepository.Create(ctx, user)
}

func (r *userRepository) GetUser(ctx context.Context, id string) (*models.User, error) {
	var entity models.User
	err := r.db.WithContext(ctx).First(&entity, "user_id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &entity, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	return r.BaseRepository.Update(ctx, user)
}

func (r *userRepository) UpdateUserFields(ctx context.Context, id string, fields map[string]any) error {
	if len(fields) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Model(&models.User{}).Where("user_id = ?", id).Updates(fields).Error
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}
