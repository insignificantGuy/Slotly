package user

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	"github.com/insignificantGuy/Slotly/internal/repository/user"
	userrole "github.com/insignificantGuy/Slotly/internal/repository/user_role"
)

const customerRoleID uint = 4

type UserService struct {
	userRepository     user.UserRepository
	userRoleRepository userrole.UserRoleRepository
}

func NewUserService(userRepository user.UserRepository, userRoleRepository userrole.UserRoleRepository) *UserService {
	return &UserService{userRepository: userRepository, userRoleRepository: userRoleRepository}
}

func (s *UserService) CreateUser(ctx context.Context, newUser *model.User) error {
	existing, err := s.userRepository.GetUserByEmail(ctx, newUser.Email)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}
	if existing != nil {
		return fmt.Errorf("user with email %s already exists", newUser.Email)
	}

	if err := s.userRepository.CreateUser(ctx, newUser); err != nil {
		return err
	}

	created, err := s.userRepository.GetUserByEmail(ctx, newUser.Email)
	if err != nil {
		return err
	}

	mapping := &model.UserRoleMapping{
		UserID: created.UserID,
		RoleID: customerRoleID,
	}
	return s.userRoleRepository.CreateUserRole(ctx, mapping)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user with id %s not found", id)
	}
	return s.userRepository.GetUser(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, req *UpdateUserRequest) error {
	existing, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return err
	}

	updates := map[string]any{}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.PhoneNumber != nil {
		updates["phone_number"] = *req.PhoneNumber
	}
	if req.Password != nil {
		updates["password"] = *req.Password
	}
	if req.FirstName != nil || req.LastName != nil {
		first, last := splitFullName(existing.FullName)
		if req.FirstName != nil {
			first = *req.FirstName
		}
		if req.LastName != nil {
			last = *req.LastName
		}
		updates["full_name"] = strings.TrimSpace(first + " " + last)
	}
	if len(updates) == 0 {
		return nil
	}

	return s.userRepository.UpdateUserFields(ctx, id, updates)
}

func splitFullName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], strings.Join(parts[1:], " ")
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	user, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("user with id %s not found", id)
	}
	user.IsDeleted = true
	user.DeletedAt = time.Now().UTC()
	return s.userRepository.UpdateUser(ctx, user)
}
