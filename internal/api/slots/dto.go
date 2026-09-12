package slots

import "time"

type CreateSlotRequest struct {
	ListingID string    `json:"listing_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
	Duration  int       `json:"duration"`
}

type UpdateSlotRequest struct {
	Date      *time.Time `json:"date"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	Duration  *int       `json:"duration"`
}
