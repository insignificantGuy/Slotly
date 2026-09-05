package userrole

import (
	"context"
	"errors"

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
	var entity models.UserRoleMapping
	err := r.db.WithContext(ctx).First(&entity, "user_id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &entity, nil
}

func (r *userRoleRepository) UpdateUserRole(ctx context.Context, userRole *models.UserRoleMapping, roleID uint) error {
	var entity models.UserRoleMapping
	err := r.db.WithContext(ctx).Model(&entity).Where("user_id = ?", userRole.UserID).Update("role_id", roleID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRoleRepository) DeleteUserRole(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}
