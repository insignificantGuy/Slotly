package listing

import "github.com/gin-gonic/gin"

type ListingController struct {
	listingsService *ListingService
}

func NewListingController(listingsService *ListingService) *ListingController {
	return &ListingController{listingsService: listingsService}
}

func (lc *ListingController) CreateListing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (lc *ListingController) GetListing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (lc *ListingController) UpdateListing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (lc *ListingController) DeleteListing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (lc *ListingController) GetListingByType(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (lc *ListingController) GetListingByCompanyID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (lc *ListingController) GetListingByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (lc *ListingController) GetListingsByPrice(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}
