package company

import (
	"errors"

	"github.com/gin-gonic/gin"
	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
)

type CompanyController struct {
	companyService *CompanyService
}

func NewCompanyController(companyService *CompanyService) *CompanyController {
	return &CompanyController{companyService: companyService}
}

func (cc *CompanyController) CreateCompany(c *gin.Context) {
	var company *model.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := cc.companyService.CreateCompany(c.Request.Context(), company); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, company)
}

func (cc *CompanyController) GetCompany(c *gin.Context) {
	var id string = c.Param("id")
	company, err := cc.companyService.GetCompany(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if company == nil {
		c.JSON(404, gin.H{"error": "company not found"})
		return
	}
}

func (cc *CompanyController) UpdateCompany(c *gin.Context) {
	id := c.Param("id")
	var company model.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	company.CompanyID = id
	if err := cc.companyService.UpdateCompany(c.Request.Context(), &company); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(404, gin.H{"error": "company not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "company updated successfully"})
}

func (cc *CompanyController) DeleteCompany(c *gin.Context) {
	var id string = c.Param("id")
	if err := cc.companyService.DeleteCompany(c.Request.Context(), id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Company deleted"})
}
