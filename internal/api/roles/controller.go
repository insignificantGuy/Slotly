package roles

import "github.com/gin-gonic/gin"

type RoleController struct {
	roleService *RoleService
}

func NewRoleController(roleService *RoleService) *RoleController {
	return &RoleController{roleService: roleService}
}

func (rc *RoleController) CreateRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (rc *RoleController) GetRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (rc *RoleController) UpdateRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (rc *RoleController) DeleteRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}
