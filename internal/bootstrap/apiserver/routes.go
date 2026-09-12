package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/insignificantGuy/Slotly/internal/auth"
)

func SetupRoutes(router *gin.Engine, tokens *auth.Tokens) {
	SetupHealthRoutes(router)
	SetupPublicRoutes(router)
	SetupProtectedRoutes(router, tokens)
}

func SetupHealthRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})
	router.GET("/live", func(c *gin.Context) {
		c.JSON(200, gin.H{"live": "OK"})
	})
}

func SetupPublicRoutes(router *gin.Engine) {
	public := router.Group("/v1")
	public.POST("/users", controller.UserController.CreateUser)
	public.POST("/auth/login", controller.AuthController.Login)
	public.POST("/auth/refresh", controller.AuthController.Refresh)
	public.POST("/auth/logout", controller.AuthController.Logout)

	public.GET("/company/:id", controller.CompanyController.GetCompany)
	public.GET("/listing/type/:type", controller.ListingController.GetListingByType)
	public.GET("/listing/company/:company_id", controller.ListingController.GetListingByCompanyID)
	public.GET("/listings/price/:price", controller.ListingController.GetListingsByPrice)
	public.GET("/listing/:id", controller.ListingController.GetListing)
	public.GET("/slot/:id", controller.SlotController.GetSlot)
	public.GET("/listing/:id/slots", controller.SlotController.GetSlotsByListingID)
}

func SetupProtectedRoutes(router *gin.Engine, tokens *auth.Tokens) {
	authed := router.Group("/v1")
	authed.Use(auth.RequireAuth(tokens))

	authed.POST("/auth/logout-all", controller.AuthController.LogoutAll)
	authed.GET("/me", controller.UserController.GetMe)
	authed.GET("/user/:id", controller.UserController.GetUser)
	authed.PUT("/user/:id", controller.UserController.UpdateUser)
	authed.DELETE("/user/:id", controller.UserController.DeleteUser)

	authed.POST("/companies", controller.CompanyController.CreateCompany)
	authed.PUT("/company/:id", controller.CompanyController.UpdateCompany)
	authed.DELETE("/company/:id", controller.CompanyController.DeleteCompany)

	authed.GET("/company-role/:id", controller.CompanyRoleController.GetCompanyRole)

	authed.POST("/listings", controller.ListingController.CreateListing)
	authed.PUT("/listing/:id", controller.ListingController.UpdateListing)
	authed.DELETE("/listing/:id", controller.ListingController.DeleteListing)

	authed.GET("/user-role", controller.UserRoleController.GetUserRole)

	authed.POST("/slots", controller.SlotController.CreateSlot)
	authed.PUT("/slot/:id", controller.SlotController.UpdateSlot)
	authed.DELETE("/slot/:id", controller.SlotController.DeleteSlot)

	authed.POST("/bookings", controller.BookingController.CreateBooking)
	authed.GET("/bookings", controller.BookingController.GetMyBookings)
	authed.GET("/bookings/listing/:listing_id", controller.BookingController.GetBookingsByListingID)
	authed.GET("/booking/:id", controller.BookingController.GetBooking)
	authed.PUT("/booking/:id", controller.BookingController.UpdateBooking)
	authed.DELETE("/booking/:id", controller.BookingController.CancelBooking)
}
