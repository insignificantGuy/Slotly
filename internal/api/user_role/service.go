package userrole

import (
	"context"
	"errors"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	userRepo "github.com/insignificantGuy/Slotly/internal/repository/user"
	userroleRepo "github.com/insignificantGuy/Slotly/internal/repository/user_role"
)

type UserRoleService struct {
	userRoleRepository userroleRepo.UserRoleRepository
	userRepository     userRepo.UserRepository
}

func NewUserRoleService(
	userRoleRepository userroleRepo.UserRoleRepository,
	userRepository userRepo.UserRepository,
) *UserRoleService {
	return &UserRoleService{
		userRoleRepository: userRoleRepository,
		userRepository:     userRepository,
	}
}

func (s *UserRoleService) CreateUserRole(ctx context.Context, userID string, roleID int) error {
	if _, err := s.userRepository.GetUser(ctx, userID); err != nil {
		return err
	}

	existing, err := s.userRoleRepository.GetUserRole(ctx, userID)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return errors.New("role already exists")
	}

	mapping := &model.UserRoleMapping{
		UserID: userID,
		RoleID: uint(roleID),
	}
	return s.userRoleRepository.CreateUserRole(ctx, mapping)
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
	if _, err := s.userRepository.GetUser(ctx, userID); err != nil {
		return err
	}

	userRole, err := s.userRoleRepository.GetUserRole(ctx, userID)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if userRole == nil || errors.Is(err, repository.ErrNotFound) {
		mapping := &model.UserRoleMapping{
			UserID: userID,
			RoleID: uint(roleID),
		}
		return s.userRoleRepository.CreateUserRole(ctx, mapping)
	}
	userRole.RoleID = uint(roleID)
	return s.userRoleRepository.UpdateUserRole(ctx, userRole, uint(roleID))
}

func (s *UserRoleService) DeleteUserRole(ctx context.Context, id string) error {
	return s.userRoleRepository.DeleteUserRole(ctx, id)
}
