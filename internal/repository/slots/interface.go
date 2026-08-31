package slots

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
)

type SlotsRepository interface {
	repository.Repository[models.Slot]
	CreateSlot(ctx context.Context, slot *models.Slot) error
	GetSlot(ctx context.Context, id string) (*models.Slot, error)
	UpdateSlot(ctx context.Context, slot *models.Slot) error
	DeleteSlot(ctx context.Context, id string) error
	GetSlotsByListingID(ctx context.Context, listingID string) (*[]models.Slot, error)
}
