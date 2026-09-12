package booking

type CreateBookingRequest struct {
	UserID    string `json:"user_id" binding:"required"`
	SlotID    uint   `json:"slot_id" binding:"required"`
	ListingID string `json:"listing_id" binding:"required"`
}

type UpdateBookingRequest struct {
	UserID string  `json:"user_id" binding:"required"`
	SlotID *uint   `json:"slot_id"`
	Status *string `json:"status"`
}

type DeleteBookingRequest struct {
	UserID string `json:"user_id" binding:"required"`
}
