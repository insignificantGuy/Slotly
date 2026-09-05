package company

import (
	"context"
	"errors"

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
	var entity model.Company
	err := r.db.WithContext(ctx).First(&entity, "company_id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &entity, nil
}

func (r *companyRepository) UpdateCompany(ctx context.Context, company *model.Company) error {
	return r.BaseRepository.Update(ctx, company)
}

func (r *companyRepository) FetchCompanies(ctx context.Context, offset, limit int) ([]model.Company, error) {
	return r.BaseRepository.List(ctx, offset, limit)
}
