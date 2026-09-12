package booking

import (
	"context"
	"errors"

	"github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrAlreadyBooked    = errors.New("slot is already booked")
	ErrAlreadyCancelled = errors.New("booking is already cancelled")
)

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) *bookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(ctx context.Context, booking *models.Booking) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var slot models.Slot
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&slot, booking.SlotID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return repository.ErrNotFound
			}
			return err
		}
		if slot.IsBooked {
			return ErrAlreadyBooked
		}
		if booking.ListingID != "" && slot.ListingID != booking.ListingID {
			return errors.New("slot does not belong to listing")
		}

		booking.ListingID = slot.ListingID
		if booking.Status == "" {
			booking.Status = models.BookingStatusConfirmed
		}
		if err := tx.Create(booking).Error; err != nil {
			return err
		}
		return tx.Model(&slot).Update("is_booked", true).Error
	})
}

func (r *bookingRepository) GetBooking(ctx context.Context, id uint) (*models.Booking, error) {
	var booking models.Booking
	err := r.db.WithContext(ctx).First(&booking, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) GetBookingBySlotID(ctx context.Context, slotID uint) (*models.Booking, error) {
	var booking models.Booking
	err := r.db.WithContext(ctx).Where("slot_id = ?", slotID).First(&booking).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) GetBookingsByUserID(ctx context.Context, userID string) ([]models.Booking, error) {
	var bookings []models.Booking
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepository) GetBookingsByListingID(ctx context.Context, listingID string) ([]models.Booking, error) {
	var bookings []models.Booking
	if err := r.db.WithContext(ctx).Where("listing_id = ?", listingID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepository) UpdateBooking(ctx context.Context, booking *models.Booking) error {
	return r.db.WithContext(ctx).Save(booking).Error
}

func (r *bookingRepository) CancelBooking(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var booking models.Booking
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&booking, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return repository.ErrNotFound
			}
			return err
		}
		if booking.Status == models.BookingStatusCancelled {
			return ErrAlreadyCancelled
		}
		if err := tx.Model(&booking).Update("status", models.BookingStatusCancelled).Error; err != nil {
			return err
		}
		return tx.Model(&models.Slot{}).Where("id = ?", booking.SlotID).Update("is_booked", false).Error
	})
}

func (r *bookingRepository) RescheduleBooking(ctx context.Context, bookingID, newSlotID uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var booking models.Booking
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&booking, bookingID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return repository.ErrNotFound
			}
			return err
		}
		if booking.Status == models.BookingStatusCancelled {
			return ErrAlreadyCancelled
		}
		if booking.SlotID == newSlotID {
			return nil
		}

		var newSlot models.Slot
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&newSlot, newSlotID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return repository.ErrNotFound
			}
			return err
		}
		if newSlot.IsBooked {
			return ErrAlreadyBooked
		}
		if newSlot.ListingID != booking.ListingID {
			return errors.New("slot does not belong to listing")
		}

		if err := tx.Model(&models.Slot{}).Where("id = ?", booking.SlotID).Update("is_booked", false).Error; err != nil {
			return err
		}
		if err := tx.Model(&newSlot).Update("is_booked", true).Error; err != nil {
			return err
		}
		return tx.Model(&booking).Update("slot_id", newSlotID).Error
	})
}
