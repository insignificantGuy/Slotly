package booking

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/insignificantGuy/Slotly/internal/repository"
	bookingRepo "github.com/insignificantGuy/Slotly/internal/repository/booking"
)

type BookingController struct {
	bookingService *BookingService
}

func NewBookingController(bookingService *BookingService) *BookingController {
	return &BookingController{bookingService: bookingService}
}

func (bc *BookingController) CreateBooking(c *gin.Context) {
	var req CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	booking, err := bc.bookingService.CreateBooking(c.Request.Context(), &req)
	if err != nil {
		bc.writeError(c, err, "listing or slot not found")
		return
	}
	c.JSON(200, gin.H{"message": "booking created successfully", "booking": booking})
}

func (bc *BookingController) GetBooking(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	booking, err := bc.bookingService.GetBooking(c.Request.Context(), id)
	if err != nil {
		bc.writeError(c, err, "booking not found")
		return
	}
	c.JSON(200, booking)
}

func (bc *BookingController) GetBookingsByUserID(c *gin.Context) {
	bookings, err := bc.bookingService.GetBookingsByUserID(c.Request.Context(), c.Param("user_id"))
	if err != nil {
		bc.writeError(c, err, "user not found")
		return
	}
	c.JSON(200, bookings)
}

func (bc *BookingController) GetBookingsByListingID(c *gin.Context) {
	userID := c.Query("user_id")
	bookings, err := bc.bookingService.GetBookingsByListingID(c.Request.Context(), c.Param("listing_id"), userID)
	if err != nil {
		bc.writeError(c, err, "listing not found")
		return
	}
	c.JSON(200, bookings)
}

func (bc *BookingController) UpdateBooking(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	var req UpdateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := bc.bookingService.UpdateBooking(c.Request.Context(), id, &req); err != nil {
		bc.writeError(c, err, "booking not found")
		return
	}
	c.JSON(200, gin.H{"message": "booking updated successfully"})
}

func (bc *BookingController) CancelBooking(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	var req DeleteBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := bc.bookingService.CancelBooking(c.Request.Context(), id, req.UserID); err != nil {
		bc.writeError(c, err, "booking not found")
		return
	}
	c.JSON(200, gin.H{"message": "booking cancelled successfully"})
}

func parseID(c *gin.Context, param string) (uint, bool) {
	id, err := strconv.ParseUint(c.Param(param), 10, 64)
	if err != nil || id == 0 {
		c.JSON(400, gin.H{"error": "invalid id"})
		return 0, false
	}
	return uint(id), true
}

func (bc *BookingController) writeError(c *gin.Context, err error, notFoundMessage string) {
	if errors.Is(err, ErrForbidden) {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, ErrAlreadyBooked) || errors.Is(err, bookingRepo.ErrAlreadyBooked) {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, ErrAlreadyCancelled) || errors.Is(err, ErrInvalidStatus) || errors.Is(err, bookingRepo.ErrAlreadyCancelled) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, ErrUserNotFound) {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, repository.ErrNotFound) {
		c.JSON(404, gin.H{"error": notFoundMessage})
		return
	}
	c.JSON(500, gin.H{"error": err.Error()})
}
