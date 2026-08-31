package user

import (
	"context"

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
	return r.BaseRepository.Get(ctx, id)
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	return r.BaseRepository.Update(ctx, user)
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}
