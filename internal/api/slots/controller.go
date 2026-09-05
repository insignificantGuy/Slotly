package slots

import "github.com/gin-gonic/gin"

type SlotController struct {
	slotService *SlotService
}

func NewSlotController(slotService *SlotService) *SlotController {
	return &SlotController{slotService: slotService}
}

func (sc *SlotController) CreateSlot(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (sc *SlotController) GetSlot(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (sc *SlotController) UpdateSlot(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (sc *SlotController) DeleteSlot(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (sc *SlotController) GetSlotsByListingID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}
