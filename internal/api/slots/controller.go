package slots

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/insignificantGuy/Slotly/internal/auth"
	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
)

type SlotController struct {
	slotService *SlotService
}

func NewSlotController(slotService *SlotService) *SlotController {
	return &SlotController{slotService: slotService}
}

func (sc *SlotController) CreateSlot(c *gin.Context) {
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	var req CreateSlotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	slot := &model.Slot{
		ListingID: req.ListingID,
		Date:      req.Date,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Duration:  req.Duration,
		IsBooked:  false,
	}
	if err := sc.slotService.CreateSlot(c.Request.Context(), slot, userID); err != nil {
		sc.writeError(c, err, "listing not found")
		return
	}
	c.JSON(200, gin.H{"message": "slot created successfully"})
}

func (sc *SlotController) GetSlot(c *gin.Context) {
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	slot, err := sc.slotService.GetSlot(c.Request.Context(), id)
	if err != nil {
		sc.writeError(c, err, "slot not found")
		return
	}
	c.JSON(200, slot)
}

func (sc *SlotController) UpdateSlot(c *gin.Context) {
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	var req UpdateSlotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := sc.slotService.UpdateSlot(c.Request.Context(), id, userID, &req); err != nil {
		sc.writeError(c, err, "slot not found")
		return
	}
	c.JSON(200, gin.H{"message": "slot updated successfully"})
}

func (sc *SlotController) DeleteSlot(c *gin.Context) {
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	id, ok := parseID(c, "id")
	if !ok {
		return
	}
	if err := sc.slotService.DeleteSlot(c.Request.Context(), id, userID); err != nil {
		sc.writeError(c, err, "slot not found")
		return
	}
	c.JSON(200, gin.H{"message": "slot deleted successfully"})
}

func (sc *SlotController) GetSlotsByListingID(c *gin.Context) {
	slots, err := sc.slotService.GetSlotsByListingID(c.Request.Context(), c.Param("id"))
	if err != nil {
		sc.writeError(c, err, "listing not found")
		return
	}
	c.JSON(200, slots)
}

func parseID(c *gin.Context, param string) (uint, bool) {
	id, err := strconv.ParseUint(c.Param(param), 10, 64)
	if err != nil || id == 0 {
		c.JSON(400, gin.H{"error": "invalid id"})
		return 0, false
	}
	return uint(id), true
}

func (sc *SlotController) writeError(c *gin.Context, err error, notFoundMessage string) {
	if errors.Is(err, ErrForbidden) {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, ErrSlotBooked) || errors.Is(err, ErrOverlap) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if errors.Is(err, repository.ErrNotFound) {
		c.JSON(404, gin.H{"error": notFoundMessage})
		return
	}
	c.JSON(500, gin.H{"error": err.Error()})
}
