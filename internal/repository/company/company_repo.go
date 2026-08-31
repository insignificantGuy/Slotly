package company

import (
	"context"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"gorm.io/gorm"
)

type companyRepository struct {
	*repository.BaseRepository[model.Company]
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *companyRepository {
	return &companyRepository{
		BaseRepository: repository.NewBaseRepository[model.Company](db),
		db:             db,
	}
}

func (r *companyRepository) RegisterCompany(ctx context.Context, company *model.Company) error {
	return r.BaseRepository.Create(ctx, company)
}

func (r *companyRepository) FetchCompanyByID(ctx context.Context, id string) (*model.Company, error) {
	return r.BaseRepository.Get(ctx, id)
}

func (r *companyRepository) UpdateCompany(ctx context.Context, company *model.Company) error {
	return r.BaseRepository.Update(ctx, company)
}

func (r *companyRepository) DeleteCompany(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}

func (r *companyRepository) FetchCompanies(ctx context.Context, offset, limit int) ([]model.Company, error) {
	return r.BaseRepository.List(ctx, offset, limit)
}
