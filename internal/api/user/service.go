package user

import (
	"context"
	"fmt"
	"strings"
	"time"

	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository/user"
)

type UserService struct {
	userRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	user, err := s.userRepository.GetUser(ctx, user.UserID)
	if err != nil {
		return err
	}
	if user != nil {
		return fmt.Errorf("user with id %s already exists", user.UserID)
	}
	return s.userRepository.CreateUser(ctx, user)
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
