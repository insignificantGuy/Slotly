package slots

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
)

type SlotsRepository interface {
	CreateSlot(ctx context.Context, slot *models.Slot) error
	GetSlotByID(ctx context.Context, id uint) (*models.Slot, error)
	GetCurrentDateSlot(ctx context.Context, listingID string) (*[]models.Slot, error)
	UpdateSlot(ctx context.Context, slot *models.Slot) error
	DeleteSlot(ctx context.Context, id uint) error
	GetSlotsByListingID(ctx context.Context, listingID string) (*[]models.Slot, error)
}
