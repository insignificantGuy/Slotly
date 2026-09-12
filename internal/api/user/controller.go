package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	authapi "github.com/insignificantGuy/Slotly/internal/api/auth"
	"github.com/insignificantGuy/Slotly/internal/auth"
	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
)

type UserController struct {
	userService *UserService
	authService *authapi.AuthService
}

func NewUserController(userService *UserService, authService *authapi.AuthService) *UserController {
	return &UserController{userService: userService, authService: authService}
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
	created, err := uc.userService.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	tokens, err := uc.authService.IssueSession(c.Request.Context(), created.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User created successfully", "user_id": created.UserID, "tokens": tokens})
}

func (uc *UserController) GetMe(c *gin.Context) {
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	uc.writeUser(c, userID)
}

func (uc *UserController) GetUser(c *gin.Context) {
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	if c.Param("id") != userID {
		c.JSON(403, gin.H{"error": "cannot view another user"})
		return
	}
	uc.writeUser(c, userID)
}

func (uc *UserController) writeUser(c *gin.Context, userID string) {
	user, err := uc.userService.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	if c.Param("id") != userID {
		c.JSON(403, gin.H{"error": "cannot update another user"})
		return
	}
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
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	if c.Param("id") != userID {
		c.JSON(403, gin.H{"error": "cannot delete another user"})
		return
	}
	err := uc.userService.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "user deleted successfully"})
}
