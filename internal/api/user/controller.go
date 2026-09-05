package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
)

type UserController struct {
	userService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		c.JSON(400, gin.H{"error": bindErr.Error()})
		return
	}
	user := model.User{
		FullName:    (req.FirstName + " " + req.LastName),
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
	}
	if err := uc.userService.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User created successfully"})
}

func (uc *UserController) GetUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := uc.userService.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var req UpdateUserRequest
	if bindErr := c.ShouldBindJSON(&req); bindErr != nil {
		c.JSON(400, gin.H{"error": bindErr.Error()})
		return
	}
	err := uc.userService.UpdateUser(c.Request.Context(), userID, &req)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user updated successfully"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	err := uc.userService.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user deleted successfully"})
}
