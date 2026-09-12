package company

type CreateCompanyRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
}
