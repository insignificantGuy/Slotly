package userrole

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/insignificantGuy/Slotly/internal/auth"
	"github.com/insignificantGuy/Slotly/internal/repository"
)

type UserRoleController struct {
	userRoleService *UserRoleService
}

func NewUserRoleController(userRoleService *UserRoleService) *UserRoleController {
	return &UserRoleController{userRoleService: userRoleService}
}

func (urc *UserRoleController) GetUserRole(c *gin.Context) {
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	userRole, err := urc.userRoleService.GetUserRole(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(404, gin.H{"error": "user role not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if userRole == nil {
		c.JSON(404, gin.H{"error": "user role not found"})
		return
	}
	c.JSON(200, userRole)
}
