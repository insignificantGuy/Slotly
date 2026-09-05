package companyrole

import "github.com/gin-gonic/gin"

type CompanyRoleController struct {
	companyRoleService *CompanyRoleService
}

func NewCompanyRoleController(companyRoleService *CompanyRoleService) *CompanyRoleController {
	return &CompanyRoleController{companyRoleService: companyRoleService}
}

func (crc *CompanyRoleController) CreateCompanyRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (crc *CompanyRoleController) GetCompanyRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (crc *CompanyRoleController) UpdateCompanyRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

func (crc *CompanyRoleController) DeleteCompanyRole(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}
