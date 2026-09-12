package api

import (
	authapi "github.com/insignificantGuy/Slotly/internal/api/auth"
	"github.com/insignificantGuy/Slotly/internal/api/booking"
	"github.com/insignificantGuy/Slotly/internal/api/company"
	companyrole "github.com/insignificantGuy/Slotly/internal/api/company_role"
	"github.com/insignificantGuy/Slotly/internal/api/listing"
	"github.com/insignificantGuy/Slotly/internal/api/slots"
	"github.com/insignificantGuy/Slotly/internal/api/user"
	userrole "github.com/insignificantGuy/Slotly/internal/api/user_role"
)

type Controller struct {
	AuthController        *authapi.AuthController
	UserController        *user.UserController
	CompanyController     *company.CompanyController
	CompanyRoleController *companyrole.CompanyRoleController
	ListingController     *listing.ListingController
	UserRoleController    *userrole.UserRoleController
	SlotController        *slots.SlotController
	BookingController     *booking.BookingController
}

func NewController(svc *Service) *Controller {
	return &Controller{
		AuthController:        authapi.NewAuthController(svc.AuthService),
		UserController:        user.NewUserController(svc.UserService, svc.AuthService),
		CompanyController:     company.NewCompanyController(svc.CompanyService, svc.CompanyRoleService),
		CompanyRoleController: companyrole.NewCompanyRoleController(svc.CompanyRoleService),
		ListingController:     listing.NewListingController(svc.ListingService),
		UserRoleController:    userrole.NewUserRoleController(svc.UserRoleService),
		SlotController:        slots.NewSlotController(svc.SlotService),
		BookingController:     booking.NewBookingController(svc.BookingService),
	}
}
