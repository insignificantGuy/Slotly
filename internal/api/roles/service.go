package roles

import (
	"context"

	model "github.com/insignificantGuy/Slotly/internal/models"
	rolesRepo "github.com/insignificantGuy/Slotly/internal/repository/roles"
)

type RoleService struct {
	rolesRepository rolesRepo.RolesRepository
}

func NewRoleService(rolesRepository rolesRepo.RolesRepository) *RoleService {
	return &RoleService{rolesRepository: rolesRepository}
}

func (s *RoleService) CreateRole(ctx context.Context, role *model.Role) error {
	return s.rolesRepository.CreateRole(ctx, role)
}

func (s *RoleService) GetRole(ctx context.Context, id string) (*model.Role, error) {
	return s.rolesRepository.GetRole(ctx, id)
}

func (s *RoleService) UpdateRole(ctx context.Context, role *model.Role) error {
	return s.rolesRepository.UpdateRole(ctx, role)
}

func (s *RoleService) DeleteRole(ctx context.Context, id string) error {
	return s.rolesRepository.DeleteRole(ctx, id)
}
