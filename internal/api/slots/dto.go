package slots

import "time"

type CreateSlotRequest struct {
	UserID    string    `json:"user_id" binding:"required"`
	ListingID string    `json:"listing_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
	Duration  int       `json:"duration"`
}

type UpdateSlotRequest struct {
	UserID    string     `json:"user_id" binding:"required"`
	Date      *time.Time `json:"date"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	Duration  *int       `json:"duration"`
}

type DeleteSlotRequest struct {
	UserID string `json:"user_id" binding:"required"`
}
