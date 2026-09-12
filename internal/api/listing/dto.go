package listing

type CreateListingRequest struct {
	UserID      string  `json:"user_id" binding:"required"`
	CompanyID   string  `json:"company_id" binding:"required"`
	Type        string  `json:"type" binding:"required"`
	Title       string  `json:"title" binding:"required"`
	Image       string  `json:"image" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

type UpdateListingRequest struct {
	UserID      string   `json:"user_id" binding:"required"`
	Type        *string  `json:"type"`
	Title       *string  `json:"title"`
	Image       *string  `json:"image"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
}

type DeleteListingRequest struct {
	UserID string `json:"user_id" binding:"required"`
}
