package companyrole

import (
	"github.com/gin-gonic/gin"
)

type CompanyRoleController struct {
	companyRoleService *CompanyRoleService
}

func NewCompanyRoleController(companyRoleService *CompanyRoleService) *CompanyRoleController {
	return &CompanyRoleController{companyRoleService: companyRoleService}
}

func (crc *CompanyRoleController) GetCompanyRole(c *gin.Context) {
	c.JSON(501, gin.H{"error": "not implemented"})
}
