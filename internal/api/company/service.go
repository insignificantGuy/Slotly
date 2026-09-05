package company

import (
	"context"
	"fmt"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"github.com/insignificantGuy/Slotly/internal/repository/company"
)

type CompanyService struct {
	companyRepository company.CompanyRepository
}

func NewCompanyService(companyRepository company.CompanyRepository) *CompanyService {
	return &CompanyService{companyRepository: companyRepository}
}

func (s *CompanyService) CreateCompany(ctx context.Context, company *model.Company) error {
	return s.companyRepository.RegisterCompany(ctx, company)
}

func (s *CompanyService) GetCompany(ctx context.Context, id string) (*model.Company, error) {
	company, err := s.companyRepository.FetchCompanyByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, nil
	}
	return company, nil
}

func (s *CompanyService) UpdateCompany(ctx context.Context, incoming *model.Company) error {
	existing, err := s.companyRepository.FetchCompanyByID(ctx, incoming.CompanyID)
	if err != nil {
		return err
	}
	if existing == nil {
		return repository.ErrNotFound
	}
	if incoming.Email != "" {
		existing.Email = incoming.Email
	}
	if incoming.Password != "" {
		existing.Password = incoming.Password
	}
	return s.companyRepository.UpdateCompany(ctx, existing)
}

func (s *CompanyService) DeleteCompany(ctx context.Context, id string) error {
	company, err := s.companyRepository.FetchCompanyByID(ctx, id)
	if err != nil {
		return err
	}
	if company == nil {
		return fmt.Errorf("company not found")
	}
	company.IsDeleted = true
	return s.companyRepository.UpdateCompany(ctx, company)
}
