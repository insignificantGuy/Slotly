package userrole

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"gorm.io/gorm"
)

type userRoleRepository struct {
	*repository.BaseRepository[models.UserRoleMapping]
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) *userRoleRepository {
	return &userRoleRepository{
		BaseRepository: repository.NewBaseRepository[models.UserRoleMapping](db),
		db:             db,
	}
}

func (r *userRoleRepository) CreateUserRole(ctx context.Context, userRole *models.UserRoleMapping) error {
	return r.BaseRepository.Create(ctx, userRole)
}

func (r *userRoleRepository) GetUserRole(ctx context.Context, id string) (*models.UserRoleMapping, error) {
	return r.BaseRepository.Get(ctx, id)
}

func (r *userRoleRepository) UpdateUserRole(ctx context.Context, userRole *models.UserRoleMapping) error {
	return r.BaseRepository.Update(ctx, userRole)
}

func (r *userRoleRepository) DeleteUserRole(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}
