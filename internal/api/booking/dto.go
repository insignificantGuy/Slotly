package booking

type CreateBookingRequest struct {
	SlotID    uint   `json:"slot_id" binding:"required"`
	ListingID string `json:"listing_id" binding:"required"`
}

type UpdateBookingRequest struct {
	SlotID *uint   `json:"slot_id"`
	Status *string `json:"status"`
}
