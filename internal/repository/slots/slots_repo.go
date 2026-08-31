package slots

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"gorm.io/gorm"
)

type slotsRepository struct {
	*repository.BaseRepository[models.Slot]
	db *gorm.DB
}

func NewSlotsRepository(db *gorm.DB) *slotsRepository {
	return &slotsRepository{
		BaseRepository: repository.NewBaseRepository[models.Slot](db),
		db:             db,
	}
}

func (r *slotsRepository) CreateSlot(ctx context.Context, slot *models.Slot) error {
	return r.BaseRepository.Create(ctx, slot)
}

func (r *slotsRepository) GetSlot(ctx context.Context, id string) (*models.Slot, error) {
	return r.BaseRepository.Get(ctx, id)
}

func (r *slotsRepository) UpdateSlot(ctx context.Context, slot *models.Slot) error {
	return r.BaseRepository.Update(ctx, slot)
}

func (r *slotsRepository) DeleteSlot(ctx context.Context, id string) error {
	return r.BaseRepository.Delete(ctx, id)
}

func (r *slotsRepository) GetSlotsByListingID(ctx context.Context, listingID string) (*[]models.Slot, error) {
	var slots []models.Slot
	if err := r.db.WithContext(ctx).Where("listing_id = ?", listingID).Find(&slots).Error; err != nil {
		return nil, err
	}
	return &slots, nil
}
