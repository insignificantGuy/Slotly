package userrole

import (
	"context"

	model "github.com/insignificantGuy/Slotly/internal/models"
	userroleRepo "github.com/insignificantGuy/Slotly/internal/repository/user_role"
)

type UserRoleService struct {
	userRoleRepository userroleRepo.UserRoleRepository
}

func NewUserRoleService(userRoleRepository userroleRepo.UserRoleRepository) *UserRoleService {
	return &UserRoleService{userRoleRepository: userRoleRepository}
}

func (s *UserRoleService) CreateUserRole(ctx context.Context, userID string, roleID int) error {
	var userRole *model.UserRoleMapping
	userRole.UserID = userID
	userRole.RoleID = uint(roleID)

	roles, err := s.userRoleRepository.GetUserRole(ctx, userRole.UserID)
	if err != nil {
		return err
	}

	if roles != nil {
		err = s.userRoleRepository.UpdateUserRole(ctx, roles, userRole.RoleID)
		if err != nil {
			return err
		}
		return nil
	}

	return s.userRoleRepository.CreateUserRole(ctx, userRole)
}

func (s *UserRoleService) GetUserRole(ctx context.Context, id string) (*model.UserRoleMapping, error) {
	roles, err := s.userRoleRepository.GetUserRole(ctx, id)
	if err != nil {
		return nil, err
	}

	if roles == nil {
		return nil, nil
	}

	return roles, nil
}

func (s *UserRoleService) UpdateUserRole(ctx context.Context, userID string, roleID int) error {
	userRole, err := s.userRoleRepository.GetUserRole(ctx, userID)
	if err != nil {
		return err
	}
	if userRole == nil {
		var newUserRole *model.UserRoleMapping
		newUserRole.UserID = userID
		newUserRole.RoleID = uint(roleID)
		err = s.userRoleRepository.CreateUserRole(ctx, newUserRole)
		if err != nil {
			return err
		}
		return nil
	}
	userRole.RoleID = uint(roleID)
	err = s.userRoleRepository.UpdateUserRole(ctx, userRole, uint(roleID))
	if err != nil {
		return err
	}
	return nil
}

func (s *UserRoleService) DeleteUserRole(ctx context.Context, id string) error {
	err := s.userRoleRepository.DeleteUserRole(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
