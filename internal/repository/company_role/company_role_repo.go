package companyrole

import (
	"context"
	"errors"

	"github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"gorm.io/gorm"
)

type companyRoleRepository struct {
	*repository.BaseRepository[models.CompanyRoleMapping]
	db *gorm.DB
}

func NewCompanyRoleRepository(db *gorm.DB) *companyRoleRepository {
	return &companyRoleRepository{
		BaseRepository: repository.NewBaseRepository[models.CompanyRoleMapping](db),
		db:             db,
	}
}

func (r *companyRoleRepository) CreateCompanyRole(ctx context.Context, companyRole *models.CompanyRoleMapping) error {
	return r.BaseRepository.Create(ctx, companyRole)
}

func (r *companyRoleRepository) GetCompanyRole(ctx context.Context, id string) (*models.CompanyRoleMapping, error) {
	var entity models.CompanyRoleMapping
	err := r.db.WithContext(ctx).First(&entity, "company_id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &entity, nil
}

func (r *companyRoleRepository) UpdateCompanyRole(ctx context.Context, companyRole *models.CompanyRoleMapping) error {
	return r.BaseRepository.Update(ctx, companyRole)
}
