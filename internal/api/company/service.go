package company

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"github.com/insignificantGuy/Slotly/internal/repository/company"
	companyrole "github.com/insignificantGuy/Slotly/internal/repository/company_role"
	userRepo "github.com/insignificantGuy/Slotly/internal/repository/user"
)

const orgAdminRoleID uint = 2

type CompanyService struct {
	companyRepository     company.CompanyRepository
	companyRoleRepository companyrole.CompanyRoleRepository
	userRepository        userRepo.UserRepository
}

func NewCompanyService(
	companyRepository company.CompanyRepository,
	companyRoleRepository companyrole.CompanyRoleRepository,
	userRepository userRepo.UserRepository,
) *CompanyService {
	return &CompanyService{
		companyRepository:     companyRepository,
		companyRoleRepository: companyRoleRepository,
		userRepository:        userRepository,
	}
}

func (s *CompanyService) CreateCompany(ctx context.Context, company *model.Company, userID string) error {
	if userID == "" {
		return errors.New("user id is required")
	}
	if _, err := s.userRepository.GetUser(ctx, userID); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	if company.CompanyID == "" {
		id, err := newUUID()
		if err != nil {
			return err
		}
		company.CompanyID = id
	}
	if err := s.companyRepository.RegisterCompany(ctx, company); err != nil {
		return err
	}
	mapping := &model.CompanyRoleMapping{
		UserID:    userID,
		CompanyID: company.CompanyID,
		RoleID:    orgAdminRoleID,
	}
	return s.companyRoleRepository.CreateCompanyRole(ctx, mapping)
}

func newUUID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
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
	if incoming.Name != "" {
		existing.Name = incoming.Name
	}
	if incoming.Address != "" {
		existing.Address = incoming.Address
	}
	if incoming.Phone != "" {
		existing.Phone = incoming.Phone
	}
	if incoming.Email != "" {
		existing.Email = incoming.Email
	}
	if incoming.Logo != "" {
		existing.Logo = incoming.Logo
	}
	if incoming.Description != "" {
		existing.Description = incoming.Description
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
