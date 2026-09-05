package companyrole

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
)

type CompanyRoleRepository interface {
	CreateCompanyRole(ctx context.Context, companyModel *models.CompanyRoleMapping) error
	GetCompanyRole(ctx context.Context, companyRoleID string) (*models.CompanyRoleMapping, error)
	UpdateCompanyRole(ctx context.Context, companyModel *models.CompanyRoleMapping) error
}
