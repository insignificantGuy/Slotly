package slots

import (
	"context"

	model "github.com/insignificantGuy/Slotly/internal/models"
	slotsRepo "github.com/insignificantGuy/Slotly/internal/repository/slots"
)

type SlotService struct {
	slotsRepository slotsRepo.SlotsRepository
}

func NewSlotService(slotsRepository slotsRepo.SlotsRepository) *SlotService {
	return &SlotService{slotsRepository: slotsRepository}
}

func (s *SlotService) CreateSlot(ctx context.Context, slot *model.Slot) error {
	return s.slotsRepository.CreateSlot(ctx, slot)
}

func (s *SlotService) GetSlot(ctx context.Context, id string) (*model.Slot, error) {
	return s.slotsRepository.GetSlot(ctx, id)
}

func (s *SlotService) UpdateSlot(ctx context.Context, slot *model.Slot) error {
	return s.slotsRepository.UpdateSlot(ctx, slot)
}

func (s *SlotService) DeleteSlot(ctx context.Context, id string) error {
	return s.slotsRepository.DeleteSlot(ctx, id)
}

func (s *SlotService) GetSlotsByListingID(ctx context.Context, listingID string) (*[]model.Slot, error) {
	return s.slotsRepository.GetSlotsByListingID(ctx, listingID)
}
