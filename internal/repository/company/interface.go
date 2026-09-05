package company

import (
	"context"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
)

type CompanyRepository interface {
	repository.Repository[model.Company]
	RegisterCompany(ctx context.Context, company *model.Company) error
	FetchCompanyByID(ctx context.Context, id string) (*model.Company, error)
	UpdateCompany(ctx context.Context, company *model.Company) error
}
