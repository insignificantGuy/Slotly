package listing

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
)

type ListingRepository interface {
	CreateListing(ctx context.Context, listing *models.Listing) error
	GetListing(ctx context.Context, id string) (*models.Listing, error)
	GetListingByType(ctx context.Context, listingType string) (*[]models.Listing, error)
	GetListingByCompanyID(ctx context.Context, companyID string) (*[]models.Listing, error)
	GetListingByID(ctx context.Context, id string) (*models.Listing, error)
	GetListingsByPrice(ctx context.Context, price float64, order string) (*[]models.Listing, error)
	UpdateListing(ctx context.Context, listing *models.Listing) error
	DeleteListing(ctx context.Context, id string) error
}
