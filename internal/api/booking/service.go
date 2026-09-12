package booking

import (
	"context"
	"errors"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	bookingRepo "github.com/insignificantGuy/Slotly/internal/repository/booking"
	companyroleRepo "github.com/insignificantGuy/Slotly/internal/repository/company_role"
	listingRepo "github.com/insignificantGuy/Slotly/internal/repository/listing"
	userRepo "github.com/insignificantGuy/Slotly/internal/repository/user"
)

const (
	orgAdminRoleID uint = 2
	staffRoleID    uint = 3
)

var (
	ErrForbidden        = errors.New("user is not allowed to modify this booking")
	ErrAlreadyBooked    = bookingRepo.ErrAlreadyBooked
	ErrAlreadyCancelled = bookingRepo.ErrAlreadyCancelled
	ErrUserNotFound     = errors.New("user not found")
	ErrInvalidStatus    = errors.New("status must be cancelled")
)

type BookingService struct {
	bookingRepository     bookingRepo.BookingRepository
	userRepository        userRepo.UserRepository
	listingRepository     listingRepo.ListingRepository
	companyRoleRepository companyroleRepo.CompanyRoleRepository
}

func NewBookingService(
	bookingRepository bookingRepo.BookingRepository,
	userRepository userRepo.UserRepository,
	listingRepository listingRepo.ListingRepository,
	companyRoleRepository companyroleRepo.CompanyRoleRepository,
) *BookingService {
	return &BookingService{
		bookingRepository:     bookingRepository,
		userRepository:        userRepository,
		listingRepository:     listingRepository,
		companyRoleRepository: companyRoleRepository,
	}
}

func (s *BookingService) CreateBooking(ctx context.Context, req *CreateBookingRequest) (*model.Booking, error) {
	if _, err := s.userRepository.GetUser(ctx, req.UserID); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if _, err := s.listingRepository.GetListing(ctx, req.ListingID); err != nil {
		return nil, err
	}

	booking := &model.Booking{
		UserID:    req.UserID,
		SlotID:    req.SlotID,
		ListingID: req.ListingID,
		Status:    model.BookingStatusConfirmed,
	}
	if err := s.bookingRepository.CreateBooking(ctx, booking); err != nil {
		return nil, err
	}
	return booking, nil
}

func (s *BookingService) GetBooking(ctx context.Context, id uint) (*model.Booking, error) {
	if id == 0 {
		return nil, errors.New("booking id is required")
	}
	return s.bookingRepository.GetBooking(ctx, id)
}

func (s *BookingService) GetBookingsByUserID(ctx context.Context, userID string) ([]model.Booking, error) {
	if userID == "" {
		return nil, errors.New("user id is required")
	}
	if _, err := s.userRepository.GetUser(ctx, userID); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return s.bookingRepository.GetBookingsByUserID(ctx, userID)
}

func (s *BookingService) GetBookingsByListingID(ctx context.Context, listingID, userID string) ([]model.Booking, error) {
	if userID == "" {
		return nil, errors.New("user id is required")
	}
	listing, err := s.listingRepository.GetListing(ctx, listingID)
	if err != nil {
		return nil, err
	}
	if err := s.assertCompanyMember(ctx, userID, listing.CompanyID); err != nil {
		return nil, err
	}
	return s.bookingRepository.GetBookingsByListingID(ctx, listingID)
}

func (s *BookingService) UpdateBooking(ctx context.Context, id uint, req *UpdateBookingRequest) error {
	existing, err := s.bookingRepository.GetBooking(ctx, id)
	if err != nil {
		return err
	}
	if err := s.assertCanModifyBooking(ctx, req.UserID, existing); err != nil {
		return err
	}

	if req.Status != nil {
		if *req.Status != model.BookingStatusCancelled {
			return ErrInvalidStatus
		}
		return s.bookingRepository.CancelBooking(ctx, id)
	}
	if req.SlotID != nil {
		return s.bookingRepository.RescheduleBooking(ctx, id, *req.SlotID)
	}
	return nil
}

func (s *BookingService) CancelBooking(ctx context.Context, id uint, userID string) error {
	existing, err := s.bookingRepository.GetBooking(ctx, id)
	if err != nil {
		return err
	}
	if err := s.assertCanModifyBooking(ctx, userID, existing); err != nil {
		return err
	}
	return s.bookingRepository.CancelBooking(ctx, id)
}

func (s *BookingService) assertCanModifyBooking(ctx context.Context, userID string, booking *model.Booking) error {
	if userID == "" {
		return errors.New("user id is required")
	}
	if booking.UserID == userID {
		return nil
	}
	listing, err := s.listingRepository.GetListing(ctx, booking.ListingID)
	if err != nil {
		return err
	}
	return s.assertCompanyMember(ctx, userID, listing.CompanyID)
}

func (s *BookingService) assertCompanyMember(ctx context.Context, userID, companyID string) error {
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
