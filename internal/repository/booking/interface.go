package booking

import (
	"context"

	"github.com/insignificantGuy/Slotly/internal/models"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, booking *models.Booking) error
	GetBooking(ctx context.Context, id uint) (*models.Booking, error)
	GetBookingBySlotID(ctx context.Context, slotID uint) (*models.Booking, error)
	GetBookingsByUserID(ctx context.Context, userID string) ([]models.Booking, error)
	GetBookingsByListingID(ctx context.Context, listingID string) ([]models.Booking, error)
	UpdateBooking(ctx context.Context, booking *models.Booking) error
	CancelBooking(ctx context.Context, id uint) error
	RescheduleBooking(ctx context.Context, bookingID, newSlotID uint) error
}
