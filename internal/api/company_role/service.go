package companyrole

import (
	"context"
	"errors"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	companyRepo "github.com/insignificantGuy/Slotly/internal/repository/company"
	companyroleRepo "github.com/insignificantGuy/Slotly/internal/repository/company_role"
)

type CompanyRoleService struct {
	companyRoleRepository companyroleRepo.CompanyRoleRepository
	companyRepository     companyRepo.CompanyRepository
}

func NewCompanyRoleService(
	companyRoleRepository companyroleRepo.CompanyRoleRepository,
	companyRepository companyRepo.CompanyRepository,
) *CompanyRoleService {
	return &CompanyRoleService{
		companyRoleRepository: companyRoleRepository,
		companyRepository:     companyRepository,
	}
}

func (crs *CompanyRoleService) CreateCompanyRole(ctx context.Context, companyModel *model.CompanyRoleMapping) error {
	if companyModel == nil || companyModel.CompanyID == "" {
		return errors.New("company id is required")
	}
	if companyModel.RoleID == 0 {
		companyModel.RoleID = 1
	}

	existing, err := crs.companyRoleRepository.GetCompanyRole(ctx, companyModel.CompanyID)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		existing.RoleID = companyModel.RoleID
		return crs.companyRoleRepository.UpdateCompanyRole(ctx, existing)
	}

	return crs.companyRoleRepository.CreateCompanyRole(ctx, companyModel)
}

func (crs *CompanyRoleService) GetCompanyRole(ctx context.Context, companyID string) (*model.CompanyRoleMapping, error) {
	if companyID == "" {
		return nil, errors.New("company id is required")
	}

	existing, err := crs.companyRoleRepository.GetCompanyRole(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (crs *CompanyRoleService) UpdateCompanyRole(ctx context.Context, companyModel *model.CompanyRoleMapping) error {
	if companyModel == nil || companyModel.CompanyID == "" {
		return errors.New("company id is required")
	}

	existing, err := crs.companyRoleRepository.GetCompanyRole(ctx, companyModel.CompanyID)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing == nil || errors.Is(err, repository.ErrNotFound) {
		if companyModel.RoleID == 0 {
			companyModel.RoleID = 1
		}
		return crs.companyRoleRepository.CreateCompanyRole(ctx, companyModel)
	}

	if companyModel.RoleID != 0 {
		existing.RoleID = companyModel.RoleID
	}
	return crs.companyRoleRepository.UpdateCompanyRole(ctx, existing)
}

func (crs *CompanyRoleService) DeleteCompanyRole(ctx context.Context, companyID string) error {
	if companyID == "" {
		return errors.New("company id is required")
	}

	existingCompany, err := crs.companyRepository.FetchCompanyByID(ctx, companyID)
	if err != nil {
		return err
	}
	if existingCompany == nil {
		return repository.ErrNotFound
	}
	existingCompany.IsDeleted = true
	if err := crs.companyRepository.UpdateCompany(ctx, existingCompany); err != nil {
		return err
	}
	return nil
}
