package roles

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"gorm.io/gorm"
)

type rolesRepository struct {
	*repository.BaseRepository[models.Role]
	db *gorm.DB
}

func NewRolesRepository(db *gorm.DB) *rolesRepository {
	return &rolesRepository{
		BaseRepository: repository.NewBaseRepository[models.Role](db),
		db:             db,
	}
}

func (r *rolesRepository) CreateRole(ctx context.Context, role *models.Role) error {
	return r.BaseRepository.Create(ctx, role)
}

func (r *rolesRepository) GetRole(ctx context.Context, id string) (*models.Role, error) {
	return r.BaseRepository.Get(ctx, id)
}

func (r *rolesRepository) UpdateRole(ctx context.Context, role *models.Role) error {
	return r.BaseRepository.Update(ctx, role)
}

func (r *rolesRepository) DeleteRole(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}
