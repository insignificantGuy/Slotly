package authapi

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/insignificantGuy/Slotly/internal/auth"
	model "github.com/insignificantGuy/Slotly/internal/models"
	"github.com/insignificantGuy/Slotly/internal/repository"
	refreshRepo "github.com/insignificantGuy/Slotly/internal/repository/refresh_token"
	userRepo "github.com/insignificantGuy/Slotly/internal/repository/user"
)

const refreshTTL = 7 * 24 * time.Hour

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInvalidRefresh     = errors.New("invalid or expired refresh token")
)

type AuthService struct {
	userRepository    userRepo.UserRepository
	refreshRepository refreshRepo.RefreshTokenRepository
	tokens            *auth.Tokens
}

func NewAuthService(
	userRepository userRepo.UserRepository,
	refreshRepository refreshRepo.RefreshTokenRepository,
	tokens *auth.Tokens,
) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		refreshRepository: refreshRepository,
		tokens:            tokens,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*TokenResponse, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}
	if user.IsDeleted {
		return nil, ErrInvalidCredentials
	}
	if err := auth.CheckPassword(user.Password, password); err != nil {
		return nil, ErrInvalidCredentials
	}
	return s.issueSession(ctx, user.UserID)
}

func (s *AuthService) Refresh(ctx context.Context, rawRefresh string) (*TokenResponse, error) {
	stored, err := s.refreshRepository.GetByHash(ctx, hashRefresh(rawRefresh))
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrInvalidRefresh
		}
		return nil, err
	}
	if stored.Revoked || time.Now().UTC().After(stored.ExpiresAt) {
		return nil, ErrInvalidRefresh
	}
	if err := s.refreshRepository.Revoke(ctx, stored.TokenHash); err != nil {
		return nil, err
	}
	return s.issueSession(ctx, stored.UserID)
}

func (s *AuthService) Logout(ctx context.Context, rawRefresh string) error {
	return s.refreshRepository.Revoke(ctx, hashRefresh(rawRefresh))
}

func (s *AuthService) LogoutAll(ctx context.Context, userID string) error {
	return s.refreshRepository.RevokeAllForUser(ctx, userID)
}

func (s *AuthService) IssueSession(ctx context.Context, userID string) (*TokenResponse, error) {
	return s.issueSession(ctx, userID)
}

func (s *AuthService) issueSession(ctx context.Context, userID string) (*TokenResponse, error) {
	access, _, err := s.tokens.IssueAccess(userID)
	if err != nil {
		return nil, err
	}
	rawRefresh, err := newRefreshToken()
	if err != nil {
		return nil, err
	}
	record := &model.RefreshToken{
		UserID:    userID,
		TokenHash: hashRefresh(rawRefresh),
		ExpiresAt: time.Now().UTC().Add(refreshTTL),
	}
	if err := s.refreshRepository.Create(ctx, record); err != nil {
		return nil, err
	}
	return &TokenResponse{
		AccessToken:  access,
		RefreshToken: rawRefresh,
		ExpiresIn:    int64(s.tokens.AccessTTL().Seconds()),
		TokenType:    "Bearer",
	}, nil
}

func newRefreshToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func hashRefresh(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}
