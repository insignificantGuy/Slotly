package listing

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"gorm.io/gorm"
)

type listingRepository struct {
	*repository.BaseRepository[models.Listing]
	db *gorm.DB
}

func NewListingRepository(db *gorm.DB) *listingRepository {
	return &listingRepository{
		BaseRepository: repository.NewBaseRepository[models.Listing](db),
		db:             db,
	}
}

func (r *listingRepository) CreateListing(ctx context.Context, listing *models.Listing) error {
	return r.BaseRepository.Create(ctx, listing)
}

func (r *listingRepository) GetListing(ctx context.Context, id string) (*models.Listing, error) {
	return r.BaseRepository.Get(ctx, id)
}

func (r *listingRepository) UpdateListing(ctx context.Context, listing *models.Listing) error {
	return r.BaseRepository.Update(ctx, listing)
}

func (r *listingRepository) DeleteListing(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}

func (r *listingRepository) GetListingByType(ctx context.Context, listingType string) (*[]models.Listing, error) {
	var listings []models.Listing
	if err := r.db.WithContext(ctx).Where("type = ?", listingType).Find(&listings).Error; err != nil {
		return nil, err
	}
	return &listings, nil
}

func (r *listingRepository) GetListingByCompanyID(ctx context.Context, companyID string) (*[]models.Listing, error) {
	var listings []models.Listing
	if err := r.db.WithContext(ctx).Where("company_id = ?", companyID).Find(&listings).Error; err != nil {
		return nil, err
	}
	return &listings, nil
}

func (r *listingRepository) GetListingByID(ctx context.Context, id string) (*models.Listing, error) {
	var listing models.Listing
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&listing).Error; err != nil {
		return nil, err
	}
	return &listing, nil
}

func (r *listingRepository) GetListingsByPrice(ctx context.Context, price float64, order string) (*[]models.Listing, error) {
	var listings []models.Listing
	if err := r.db.WithContext(ctx).Where("price = ?", price).Order(order).Find(&listings).Error; err != nil {
		return nil, err
	}
	return &listings, nil
}
