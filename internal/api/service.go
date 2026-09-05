package api

import (
	"github.com/insignificantGuy/Slotly/internal/api/company"
	companyrole "github.com/insignificantGuy/Slotly/internal/api/company_role"
	"github.com/insignificantGuy/Slotly/internal/api/listing"
	"github.com/insignificantGuy/Slotly/internal/api/roles"
	"github.com/insignificantGuy/Slotly/internal/api/slots"
	"github.com/insignificantGuy/Slotly/internal/api/user"
	userrole "github.com/insignificantGuy/Slotly/internal/api/user_role"
	companyRepo "github.com/insignificantGuy/Slotly/internal/repository/company"
	companyroleRepo "github.com/insignificantGuy/Slotly/internal/repository/company_role"
	listingRepo "github.com/insignificantGuy/Slotly/internal/repository/listing"
	rolesRepo "github.com/insignificantGuy/Slotly/internal/repository/roles"
	slotsRepo "github.com/insignificantGuy/Slotly/internal/repository/slots"
	userRepo "github.com/insignificantGuy/Slotly/internal/repository/user"
	userroleRepo "github.com/insignificantGuy/Slotly/internal/repository/user_role"
	"gorm.io/gorm"
)

type Service struct {
	UserService        *user.UserService
	CompanyService     *company.CompanyService
	CompanyRoleService *companyrole.CompanyRoleService
	ListingService     *listing.ListingService
	UserRoleService    *userrole.UserRoleService
	RoleService        *roles.RoleService
	SlotService        *slots.SlotService
}

func NewService(db *gorm.DB) (*Service, error) {
	return &Service{
		UserService:    user.NewUserService(userRepo.NewUserRepository(db)),
		CompanyService: company.NewCompanyService(companyRepo.NewCompanyRepository(db)),
		CompanyRoleService: companyrole.NewCompanyRoleService(
			companyroleRepo.NewCompanyRoleRepository(db),
			companyRepo.NewCompanyRepository(db),
		),
		ListingService:  listing.NewListingService(listingRepo.NewListingRepository(db)),
		UserRoleService: userrole.NewUserRoleService(userroleRepo.NewUserRoleRepository(db)),
		RoleService:     roles.NewRoleService(rolesRepo.NewRolesRepository(db)),
		SlotService:     slots.NewSlotService(slotsRepo.NewSlotsRepository(db)),
	}, nil
}
