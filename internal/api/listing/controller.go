package listing

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
)

type ListingController struct {
	listingsService *ListingService
}

func NewListingController(listingsService *ListingService) *ListingController {
	return &ListingController{listingsService: listingsService}
}

func (lc *ListingController) CreateListing(c *gin.Context) {
	var req CreateListingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	listing := model.Listing{
		Type:        req.Type,
		Title:       req.Title,
		Image:       req.Image,
		Description: req.Description,
		CompanyID:   req.CompanyID,
		Price:       req.Price,
	}
	if err := lc.listingsService.CreateListing(c.Request.Context(), &listing, req.UserID); err != nil {
		lc.writeError(c, err, "company not found")
		return
	}
	c.JSON(200, gin.H{"message": "Listing created successfully"})
}

func (lc *ListingController) GetListing(c *gin.Context) {
	listing, err := lc.listingsService.GetListing(c.Request.Context(), c.Param("id"))
	if err != nil {
		lc.writeError(c, err, "listing not found")
		return
	}
	c.JSON(200, listing)
}

func (lc *ListingController) UpdateListing(c *gin.Context) {
	var req UpdateListingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := lc.listingsService.UpdateListing(c.Request.Context(), c.Param("id"), &req); err != nil {
		lc.writeError(c, err, "listing not found")
		return
	}
	c.JSON(200, gin.H{"message": "listing updated successfully"})
}

func (lc *ListingController) DeleteListing(c *gin.Context) {
	var req DeleteListingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := lc.listingsService.DeleteListing(c.Request.Context(), c.Param("id"), req.UserID); err != nil {
		lc.writeError(c, err, "listing not found")
		return
	}
	c.JSON(200, gin.H{"message": "listing deleted successfully"})
}

func (lc *ListingController) GetListingByType(c *gin.Context) {
	listings, err := lc.listingsService.GetListingByType(c.Request.Context(), c.Param("type"))
	if err != nil {
		lc.writeError(c, err, "listing not found")
		return
	}
	c.JSON(200, listings)
}

func (lc *ListingController) GetListingByCompanyID(c *gin.Context) {
	listings, err := lc.listingsService.GetListingByCompanyID(c.Request.Context(), c.Param("company_id"))
	if err != nil {
		lc.writeError(c, err, "listing not found")
		return
	}
	c.JSON(200, listings)
}

func (lc *ListingController) GetListingByID(c *gin.Context) {
	listing, err := lc.listingsService.GetListingByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		lc.writeError(c, err, "listing not found")
		return
	}
	c.JSON(200, listing)
}

func (lc *ListingController) GetListingsByPrice(c *gin.Context) {
	price, err := strconv.ParseFloat(c.Param("price"), 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid price"})
		return
	}
	listings, err := lc.listingsService.GetListingsByPrice(c.Request.Context(), price, c.Query("order"))
	if err != nil {
		lc.writeError(c, err, "listing not found")
		return
	}
	c.JSON(200, listings)
}

func (lc *ListingController) writeError(c *gin.Context, err error, notFoundMessage string) {
	if errors.Is(err, ErrForbidden) {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, repository.ErrNotFound) {
		c.JSON(404, gin.H{"error": notFoundMessage})
		return
	}
	c.JSON(500, gin.H{"error": err.Error()})
}
