package companyrole

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
)

type CompanyRoleRepository interface {
	CreateCompanyRole(ctx context.Context, companyRole *models.CompanyRoleMapping) error
	GetCompanyRole(ctx context.Context, id string) (*models.CompanyRoleMapping, error)
	UpdateCompanyRole(ctx context.Context, companyRole *models.CompanyRoleMapping) error
	DeleteCompanyRole(ctx context.Context, id string) error
}
