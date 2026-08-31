package companyrole

import (
	"context"

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
	return r.BaseRepository.Get(ctx, id)
}

func (r *companyRoleRepository) UpdateCompanyRole(ctx context.Context, companyRole *models.CompanyRoleMapping) error {
	return r.BaseRepository.Update(ctx, companyRole)
}

func (r *companyRoleRepository) DeleteCompanyRole(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}
