package slots

import (
	"context"
	"errors"
	"time"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	companyroleRepo "github.com/insignificantGuy/Slotly/internal/repository/company_role"
	listingRepo "github.com/insignificantGuy/Slotly/internal/repository/listing"
	slotsRepo "github.com/insignificantGuy/Slotly/internal/repository/slots"
)

const (
	orgAdminRoleID uint = 2
	staffRoleID    uint = 3
)

var (
	ErrForbidden  = errors.New("user is not allowed to modify slots for this company")
	ErrSlotBooked = errors.New("cannot modify a booked slot")
	ErrOverlap    = errors.New("slot overlaps with an existing slot")
)

type SlotService struct {
	slotsRepository       slotsRepo.SlotsRepository
	listingRepository     listingRepo.ListingRepository
	companyRoleRepository companyroleRepo.CompanyRoleRepository
}

func NewSlotService(
	slotsRepository slotsRepo.SlotsRepository,
	listingRepository listingRepo.ListingRepository,
	companyRoleRepository companyroleRepo.CompanyRoleRepository,
) *SlotService {
	return &SlotService{
		slotsRepository:       slotsRepository,
		listingRepository:     listingRepository,
		companyRoleRepository: companyRoleRepository,
	}
}

func (s *SlotService) CreateSlot(ctx context.Context, slot *model.Slot, userID string) error {
	if err := s.validateTimes(slot.StartTime, slot.EndTime); err != nil {
		return err
	}
	if _, err := s.listingForWrite(ctx, slot.ListingID, userID); err != nil {
		return err
	}

	if slot.Duration == 0 {
		slot.Duration = int(slot.EndTime.Sub(slot.StartTime).Minutes())
	}
	if err := s.assertNoOverlap(ctx, slot.ListingID, 0, slot.StartTime, slot.EndTime); err != nil {
		return err
	}
	return s.slotsRepository.CreateSlot(ctx, slot)
}

func (s *SlotService) GetSlot(ctx context.Context, id uint) (*model.Slot, error) {
	if id == 0 {
		return nil, errors.New("slot id is required")
	}
	return s.slotsRepository.GetSlotByID(ctx, id)
}

func (s *SlotService) UpdateSlot(ctx context.Context, id uint, req *UpdateSlotRequest) error {
	if id == 0 {
		return errors.New("slot id is required")
	}
	existing, err := s.slotsRepository.GetSlotByID(ctx, id)
	if err != nil {
		return err
	}
	if existing.IsBooked {
		return ErrSlotBooked
	}
	if _, err := s.listingForWrite(ctx, existing.ListingID, req.UserID); err != nil {
		return err
	}

	if req.Date != nil {
		existing.Date = *req.Date
	}
	if req.StartTime != nil {
		existing.StartTime = *req.StartTime
	}
	if req.EndTime != nil {
		existing.EndTime = *req.EndTime
	}
	if req.Duration != nil {
		existing.Duration = *req.Duration
	}
	if err := s.validateTimes(existing.StartTime, existing.EndTime); err != nil {
		return err
	}
	if req.Duration == nil {
		existing.Duration = int(existing.EndTime.Sub(existing.StartTime).Minutes())
	}
	if err := s.assertNoOverlap(ctx, existing.ListingID, existing.ID, existing.StartTime, existing.EndTime); err != nil {
		return err
	}
	return s.slotsRepository.UpdateSlot(ctx, existing)
}

func (s *SlotService) DeleteSlot(ctx context.Context, id uint, userID string) error {
	if id == 0 {
		return errors.New("slot id is required")
	}
	existing, err := s.slotsRepository.GetSlotByID(ctx, id)
	if err != nil {
		return err
	}
	if existing.IsBooked {
		return ErrSlotBooked
	}
	if _, err := s.listingForWrite(ctx, existing.ListingID, userID); err != nil {
		return err
	}
	return s.slotsRepository.DeleteSlot(ctx, id)
}

func (s *SlotService) GetSlotsByListingID(ctx context.Context, listingID string) (*[]model.Slot, error) {
	if listingID == "" {
		return nil, errors.New("listing id is required")
	}
	if _, err := s.listingRepository.GetListing(ctx, listingID); err != nil {
		return nil, err
	}
	return s.slotsRepository.GetSlotsByListingID(ctx, listingID)
}

func (s *SlotService) listingForWrite(ctx context.Context, listingID, userID string) (*model.Listing, error) {
	if userID == "" {
		return nil, errors.New("user id is required")
	}
	if listingID == "" {
		return nil, errors.New("listing id is required")
	}
	listing, err := s.listingRepository.GetListing(ctx, listingID)
	if err != nil {
		return nil, err
	}
	if err := s.assertCanWriteSlots(ctx, userID, listing.CompanyID); err != nil {
		return nil, err
	}
	return listing, nil
}

func (s *SlotService) assertCanWriteSlots(ctx context.Context, userID, companyID string) error {
	membership, err := s.companyRoleRepository.GetMembership(ctx, userID, companyID)
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

func (s *SlotService) assertNoOverlap(ctx context.Context, listingID string, excludeID uint, start, end time.Time) error {
	existingSlots, err := s.slotsRepository.GetSlotsByListingID(ctx, listingID)
	if err != nil {
		return err
	}
	for _, existing := range *existingSlots {
		if excludeID != 0 && existing.ID == excludeID {
			continue
		}
		if start.Before(existing.EndTime) && existing.StartTime.Before(end) {
			return ErrOverlap
		}
	}
	return nil
}

func (s *SlotService) validateTimes(start, end time.Time) error {
	if !end.After(start) {
		return errors.New("end time must be after start time")
	}
	return nil
}
