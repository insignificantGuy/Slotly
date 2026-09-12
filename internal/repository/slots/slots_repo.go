package slots

import (
	"context"
	"errors"
	"time"

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

func (r *slotsRepository) GetSlotByID(ctx context.Context, id uint) (*models.Slot, error) {
	var slot models.Slot
	err := r.db.WithContext(ctx).First(&slot, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &slot, nil
}

func (r *slotsRepository) GetCurrentDateSlot(ctx context.Context, listingID string) (*[]models.Slot, error) {
	var slots []models.Slot
	today := time.Now().UTC().Format("2006-01-02")
	if err := r.db.WithContext(ctx).Where("listing_id = ? AND DATE(date) = ?", listingID, today).Find(&slots).Error; err != nil {
		return nil, err
	}
	return &slots, nil
}

func (r *slotsRepository) UpdateSlot(ctx context.Context, slot *models.Slot) error {
	return r.BaseRepository.Update(ctx, slot)
}

func (r *slotsRepository) DeleteSlot(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Slot{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *slotsRepository) GetSlotsByListingID(ctx context.Context, listingID string) (*[]models.Slot, error) {
	var slots []models.Slot
	if err := r.db.WithContext(ctx).Where("listing_id = ?", listingID).Find(&slots).Error; err != nil {
		return nil, err
	}
	return &slots, nil
}
