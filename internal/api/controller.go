package api

import (
	"github.com/insignificantGuy/Slotly/internal/api/booking"
	"github.com/insignificantGuy/Slotly/internal/api/company"
	companyrole "github.com/insignificantGuy/Slotly/internal/api/company_role"
	"github.com/insignificantGuy/Slotly/internal/api/listing"
	"github.com/insignificantGuy/Slotly/internal/api/roles"
	"github.com/insignificantGuy/Slotly/internal/api/slots"
	"github.com/insignificantGuy/Slotly/internal/api/user"
	userrole "github.com/insignificantGuy/Slotly/internal/api/user_role"
)

type Controller struct {
	UserController        *user.UserController
	CompanyController     *company.CompanyController
	CompanyRoleController *companyrole.CompanyRoleController
	ListingController     *listing.ListingController
	UserRoleController    *userrole.UserRoleController
	RoleController        *roles.RoleController
	SlotController        *slots.SlotController
	BookingController     *booking.BookingController
}

func NewController(svc *Service) *Controller {
	return &Controller{
		UserController:        user.NewUserController(svc.UserService),
		CompanyController:     company.NewCompanyController(svc.CompanyService, svc.CompanyRoleService),
		CompanyRoleController: companyrole.NewCompanyRoleController(svc.CompanyRoleService),
		ListingController:     listing.NewListingController(svc.ListingService),
		UserRoleController:    userrole.NewUserRoleController(svc.UserRoleService),
		RoleController:        roles.NewRoleController(svc.RoleService),
		SlotController:        slots.NewSlotController(svc.SlotService),
		BookingController:     booking.NewBookingController(svc.BookingService),
	}
}
