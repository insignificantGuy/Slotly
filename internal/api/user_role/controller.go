package userrole

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/insignificantGuy/Slotly/internal/repository"
)

type UserRoleController struct {
	userRoleService *UserRoleService
}

func NewUserRoleController(userRoleService *UserRoleService) *UserRoleController {
	return &UserRoleController{userRoleService: userRoleService}
}

func (urc *UserRoleController) CreateUserRole(c *gin.Context) {
	userID := c.Param("user_id")
	if err := urc.userRoleService.CreateUserRole(c.Request.Context(), userID, 4); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user role created successfully"})
}

func (urc *UserRoleController) GetUserRole(c *gin.Context) {
	userID := c.Param("user_id")
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

func (urc *UserRoleController) UpdateUserRole(c *gin.Context) {
	userID := c.Param("user_id")
	roleID, err := strconv.Atoi(c.Param("role_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := urc.userRoleService.UpdateUserRole(c.Request.Context(), userID, roleID); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user role updated successfully"})
}

func (urc *UserRoleController) DeleteUserRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}
