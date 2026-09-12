package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("invalid token")

type Tokens struct {
	secret []byte
	ttl    time.Duration
}

func NewTokens(secret string, ttl time.Duration) *Tokens {
	return &Tokens{secret: []byte(secret), ttl: ttl}
}

func (t *Tokens) IssueAccess(userID string) (string, time.Time, error) {
	expires := time.Now().UTC().Add(t.ttl)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"typ": "access",
		"exp": expires.Unix(),
		"iat": time.Now().UTC().Unix(),
	})
	signed, err := token.SignedString(t.secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, expires, nil
}

func (t *Tokens) ParseAccess(tokenString string) (string, error) {
	parsed, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return t.secret, nil
	})
	if err != nil || !parsed.Valid {
		return "", ErrInvalidToken
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrInvalidToken
	}
	if typ, _ := claims["typ"].(string); typ != "access" {
		return "", ErrInvalidToken
	}
	sub, _ := claims["sub"].(string)
	if sub == "" {
		return "", ErrInvalidToken
	}
	return sub, nil
}

func (t *Tokens) AccessTTL() time.Duration {
	return t.ttl
}
