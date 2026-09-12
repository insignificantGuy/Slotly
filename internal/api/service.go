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
	"github.com/insignificantGuy/Slotly/internal/auth"
	bookingRepo "github.com/insignificantGuy/Slotly/internal/repository/booking"
	companyRepo "github.com/insignificantGuy/Slotly/internal/repository/company"
	companyroleRepo "github.com/insignificantGuy/Slotly/internal/repository/company_role"
	listingRepo "github.com/insignificantGuy/Slotly/internal/repository/listing"
	refreshRepo "github.com/insignificantGuy/Slotly/internal/repository/refresh_token"
	slotsRepo "github.com/insignificantGuy/Slotly/internal/repository/slots"
	userRepo "github.com/insignificantGuy/Slotly/internal/repository/user"
	userroleRepo "github.com/insignificantGuy/Slotly/internal/repository/user_role"
	"gorm.io/gorm"
)

type Service struct {
	AuthService        *authapi.AuthService
	UserService        *user.UserService
	CompanyService     *company.CompanyService
	CompanyRoleService *companyrole.CompanyRoleService
	ListingService     *listing.ListingService
	UserRoleService    *userrole.UserRoleService
	SlotService        *slots.SlotService
	BookingService     *booking.BookingService
}

func NewService(db *gorm.DB, tokens *auth.Tokens) (*Service, error) {
	users := userRepo.NewUserRepository(db)
	userRoles := userroleRepo.NewUserRoleRepository(db)
	return &Service{
		AuthService: authapi.NewAuthService(
			users,
			refreshRepo.NewRefreshTokenRepository(db),
			tokens,
		),
		UserService: user.NewUserService(users, userRoles),
		CompanyService: company.NewCompanyService(
			companyRepo.NewCompanyRepository(db),
			companyroleRepo.NewCompanyRoleRepository(db),
			users,
		),
		CompanyRoleService: companyrole.NewCompanyRoleService(
			companyroleRepo.NewCompanyRoleRepository(db),
			companyRepo.NewCompanyRepository(db),
		),
		ListingService: listing.NewListingService(
			listingRepo.NewListingRepository(db),
			companyroleRepo.NewCompanyRoleRepository(db),
			companyRepo.NewCompanyRepository(db),
		),
		UserRoleService: userrole.NewUserRoleService(userRoles, users),
		SlotService: slots.NewSlotService(
			slotsRepo.NewSlotsRepository(db),
			listingRepo.NewListingRepository(db),
			companyroleRepo.NewCompanyRoleRepository(db),
		),
		BookingService: booking.NewBookingService(
			bookingRepo.NewBookingRepository(db),
			users,
			listingRepo.NewListingRepository(db),
			companyroleRepo.NewCompanyRoleRepository(db),
		),
	}, nil
}
