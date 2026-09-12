package apiserver

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	SetupHealthRoutes(router)
	SetupUserRoutes(router)
	SetupCompanyRoutes(router)
	SetupCompanyRoleRoutes(router)
	SetupListingRoutes(router)
	SetupUserRoleRoutes(router)
	SetupRoleRoutes(router)
	SetupSlotRoutes(router)
}

func SetupHealthRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})
	router.GET("/live", func(c *gin.Context) {
		c.JSON(200, gin.H{"live": "OK"})
	})
}

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/v1")
	userController := controller.UserController
	userGroup.POST("/users", userController.CreateUser)
	userGroup.GET("/user/:id", userController.GetUser)
	userGroup.PUT("/user/:id", userController.UpdateUser)
	userGroup.DELETE("/user/:id", userController.DeleteUser)
}

func SetupCompanyRoutes(router *gin.Engine) {
	companyGroup := router.Group("/v1")
	companyController := controller.CompanyController
	companyGroup.POST("/companies", companyController.CreateCompany)
	companyGroup.GET("/company/:id", companyController.GetCompany)
	companyGroup.PUT("/company/:id", companyController.UpdateCompany)
	companyGroup.DELETE("/company/:id", companyController.DeleteCompany)
}

func SetupCompanyRoleRoutes(router *gin.Engine) {
	companyRoleGroup := router.Group("/v1")
	companyRoleController := controller.CompanyRoleController
	companyRoleGroup.POST("/company-roles", companyRoleController.CreateCompanyRole)
	companyRoleGroup.GET("/company-role/:id", companyRoleController.GetCompanyRole)
	companyRoleGroup.PUT("/company-role/:id", companyRoleController.UpdateCompanyRole)
	companyRoleGroup.DELETE("/company-role/:id", companyRoleController.DeleteCompanyRole)
}

func SetupListingRoutes(router *gin.Engine) {
	listingGroup := router.Group("/v1")
	listingController := controller.ListingController
	listingGroup.POST("/listings", listingController.CreateListing)
	listingGroup.GET("/listing/:id", listingController.GetListing)
	listingGroup.PUT("/listing/:id", listingController.UpdateListing)
	listingGroup.DELETE("/listing/:id", listingController.DeleteListing)
	listingGroup.GET("/listing/type/:type", listingController.GetListingByType)
	listingGroup.GET("/listing/company/:company_id", listingController.GetListingByCompanyID)
	listingGroup.GET("/listing/:id", listingController.GetListingByID)
	listingGroup.GET("/listings/price/:price", listingController.GetListingsByPrice)
}

func SetupUserRoleRoutes(router *gin.Engine) {
	userRoleGroup := router.Group("/v1")
	userRoleController := controller.UserRoleController
	userRoleGroup.POST("/user-roles", userRoleController.CreateUserRole)
	userRoleGroup.GET("/user-role/:id", userRoleController.GetUserRole)
	userRoleGroup.PUT("/user-role/:id", userRoleController.UpdateUserRole)
	userRoleGroup.DELETE("/user-role/:id", userRoleController.DeleteUserRole)
}

func SetupRoleRoutes(router *gin.Engine) {
	roleGroup := router.Group("/v1")
	roleController := controller.RoleController
	roleGroup.POST("/roles", roleController.CreateRole)
	roleGroup.GET("/role/:id", roleController.GetRole)
	roleGroup.PUT("/role/:id", roleController.UpdateRole)
	roleGroup.DELETE("/role/:id", roleController.DeleteRole)
}

func SetupSlotRoutes(router *gin.Engine) {
	slotGroup := router.Group("/v1")
	slotController := controller.SlotController
	slotGroup.POST("/slots", slotController.CreateSlot)
	slotGroup.GET("/slot/:id", slotController.GetSlot)
	slotGroup.PUT("/slot/:id", slotController.UpdateSlot)
	slotGroup.DELETE("/slot/:id", slotController.DeleteSlot)
	slotGroup.GET("/listing/:id/slots", slotController.GetSlotsByListingID)
}
