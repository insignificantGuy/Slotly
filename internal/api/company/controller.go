package company

import (
	"errors"

	"github.com/gin-gonic/gin"
	companyrole "github.com/insignificantGuy/Slotly/internal/api/company_role"
	"github.com/insignificantGuy/Slotly/internal/auth"
	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
)

type CompanyController struct {
	companyService     *CompanyService
	CompanyRoleService *companyrole.CompanyRoleService
}

func NewCompanyController(companyService *CompanyService, companyRoleService *companyrole.CompanyRoleService) *CompanyController {
	return &CompanyController{companyService: companyService, CompanyRoleService: companyRoleService}
}

func (cc *CompanyController) CreateCompany(c *gin.Context) {
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	var req CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	company := model.Company{
		Name:        req.Name,
		Address:     req.Address,
		Phone:       req.Phone,
		Email:       req.Email,
		Logo:        req.Logo,
		Description: req.Description,
	}
	if err := cc.companyService.CreateCompany(c.Request.Context(), &company, userID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Company created successfully"})
}

func (cc *CompanyController) GetCompany(c *gin.Context) {
	company, err := cc.companyService.GetCompany(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if company == nil {
		c.JSON(404, gin.H{"error": "company not found"})
		return
	}
	c.JSON(200, company)
}

func (cc *CompanyController) UpdateCompany(c *gin.Context) {
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	id := c.Param("id")
	var company model.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	company.CompanyID = id
	if err := cc.companyService.UpdateCompany(c.Request.Context(), userID, &company); err != nil {
		if errors.Is(err, ErrForbidden) {
			c.JSON(403, gin.H{"error": err.Error()})
			return
		}
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
	userID, ok := auth.MustUserID(c)
	if !ok {
		return
	}
	if err := cc.companyService.DeleteCompany(c.Request.Context(), c.Param("id"), userID); err != nil {
		if errors.Is(err, ErrForbidden) {
			c.JSON(403, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Company deleted"})
}
