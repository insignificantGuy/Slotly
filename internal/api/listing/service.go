package listing

import (
	"context"

	model "github.com/insignificantGuy/Slotly/internal/models"
	listingRepo "github.com/insignificantGuy/Slotly/internal/repository/listing"
)

type ListingService struct {
	listingsRepository listingRepo.ListingRepository
}

func NewListingService(listingsRepository listingRepo.ListingRepository) *ListingService {
	return &ListingService{listingsRepository: listingsRepository}
}

func (ls *ListingService) CreateListing(ctx context.Context, listing *model.Listing) error {
	return ls.listingsRepository.CreateListing(ctx, listing)
}

func (ls *ListingService) GetListing(ctx context.Context, listingID string) (*model.Listing, error) {
	return ls.listingsRepository.GetListing(ctx, listingID)
}

func (ls *ListingService) UpdateListing(ctx context.Context, listing *model.Listing) error {
	return ls.listingsRepository.UpdateListing(ctx, listing)
}

func (ls *ListingService) DeleteListing(ctx context.Context, listingID string) error {
	return ls.listingsRepository.DeleteListing(ctx, listingID)
}

func (ls *ListingService) GetListingByType(ctx context.Context, listingType string) (*[]model.Listing, error) {
	return ls.listingsRepository.GetListingByType(ctx, listingType)
}

func (ls *ListingService) GetListingByCompanyID(ctx context.Context, companyID string) (*[]model.Listing, error) {
	return ls.listingsRepository.GetListingByCompanyID(ctx, companyID)
}

func (ls *ListingService) GetListingByID(ctx context.Context, listingID string) (*model.Listing, error) {
	return ls.listingsRepository.GetListingByID(ctx, listingID)
}

func (ls *ListingService) GetListingsByPrice(ctx context.Context, price float64, order string) (*[]model.Listing, error) {
	return ls.listingsRepository.GetListingsByPrice(ctx, price, order)
}
