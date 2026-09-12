package listing

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	companyRepo "github.com/insignificantGuy/Slotly/internal/repository/company"
	companyroleRepo "github.com/insignificantGuy/Slotly/internal/repository/company_role"
	listingRepo "github.com/insignificantGuy/Slotly/internal/repository/listing"
)

const (
	orgAdminRoleID uint = 2
	staffRoleID    uint = 3
)

var ErrForbidden = errors.New("user is not allowed to modify listings for this company")

type ListingService struct {
	listingsRepository    listingRepo.ListingRepository
	companyRoleRepository companyroleRepo.CompanyRoleRepository
	companyRepository     companyRepo.CompanyRepository
}

func NewListingService(
	listingsRepository listingRepo.ListingRepository,
	companyRoleRepository companyroleRepo.CompanyRoleRepository,
	companyRepository companyRepo.CompanyRepository,
) *ListingService {
	return &ListingService{
		listingsRepository:    listingsRepository,
		companyRoleRepository: companyRoleRepository,
		companyRepository:     companyRepository,
	}
}

func (ls *ListingService) CreateListing(ctx context.Context, listing *model.Listing, userID string) error {
	if userID == "" {
		return errors.New("user id is required")
	}
	if listing.CompanyID == "" {
		return errors.New("company id is required")
	}

	company, err := ls.companyRepository.FetchCompanyByID(ctx, listing.CompanyID)
	if err != nil {
		return err
	}
	if company == nil || company.IsDeleted {
		return repository.ErrNotFound
	}

	if err := ls.assertCanWriteListing(ctx, userID, listing.CompanyID); err != nil {
		return err
	}

	if listing.ListingID == "" {
		id, err := newUUID()
		if err != nil {
			return err
		}
		listing.ListingID = id
	}
	return ls.listingsRepository.CreateListing(ctx, listing)
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

func (ls *ListingService) assertCanWriteListing(ctx context.Context, userID, companyID string) error {
	membership, err := ls.companyRoleRepository.GetMembership(ctx, userID, companyID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrForbidden
		}
		return err
	}
	if membership.RoleID != orgAdminRoleID && membership.RoleID != staffRoleID {
		return ErrForbidden
	}
	return nil
}

func (ls *ListingService) GetListing(ctx context.Context, listingID string) (*model.Listing, error) {
	if listingID == "" {
		return nil, errors.New("listing id is required")
	}
	return ls.listingsRepository.GetListing(ctx, listingID)
}

func (ls *ListingService) UpdateListing(ctx context.Context, listingID, userID string, req *UpdateListingRequest) error {
	if listingID == "" {
		return errors.New("listing id is required")
	}
	if userID == "" {
		return errors.New("user id is required")
	}

	existing, err := ls.listingsRepository.GetListing(ctx, listingID)
	if err != nil {
		return err
	}
	if err := ls.assertCanWriteListing(ctx, userID, existing.CompanyID); err != nil {
		return err
	}

	if req.Type != nil {
		existing.Type = *req.Type
	}
	if req.Title != nil {
		existing.Title = *req.Title
	}
	if req.Image != nil {
		existing.Image = *req.Image
	}
	if req.Description != nil {
		existing.Description = *req.Description
	}
	if req.Price != nil {
		existing.Price = *req.Price
	}
	return ls.listingsRepository.UpdateListing(ctx, existing)
}

func (ls *ListingService) DeleteListing(ctx context.Context, listingID, userID string) error {
	if listingID == "" {
		return errors.New("listing id is required")
	}
	if userID == "" {
		return errors.New("user id is required")
	}

	existing, err := ls.listingsRepository.GetListing(ctx, listingID)
	if err != nil {
		return err
	}
	if err := ls.assertCanWriteListing(ctx, userID, existing.CompanyID); err != nil {
		return err
	}
	return ls.listingsRepository.DeleteListing(ctx, listingID)
}

func (ls *ListingService) GetListingByType(ctx context.Context, listingType string) (*[]model.Listing, error) {
	if listingType == "" {
		return nil, errors.New("listing type is required")
	}
	return ls.listingsRepository.GetListingByType(ctx, listingType)
}

func (ls *ListingService) GetListingByCompanyID(ctx context.Context, companyID string) (*[]model.Listing, error) {
	if companyID == "" {
		return nil, errors.New("company id is required")
	}
	return ls.listingsRepository.GetListingByCompanyID(ctx, companyID)
}

func (ls *ListingService) GetListingByID(ctx context.Context, listingID string) (*model.Listing, error) {
	return ls.GetListing(ctx, listingID)
}

func (ls *ListingService) GetListingsByPrice(ctx context.Context, price float64, order string) (*[]model.Listing, error) {
	if order != "desc" {
		order = "asc"
	}
	return ls.listingsRepository.GetListingsByPrice(ctx, price, "price "+order)
}
