package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssueAt   time.Time `json:"issue_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims
}

// Implement the GetAudience method for the Payload struct
func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	// Example logic to retrieve audience(s) - can be more complex
	if payload.Username == "" {
		return nil, fmt.Errorf("audience is missing")
	}

	// Return a sample audience as a ClaimStrings (slice of strings)
	return jwt.ClaimStrings{payload.Username}, nil
}

// GetExpirationTime implements jwt.Claims.
func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	if payload.ExpiredAt.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}
	return payload.ExpiresAt, nil
}

// GetIssuedAt implements jwt.Claims.
func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	if payload.IssueAt.After(payload.ExpiredAt) {
		return nil, fmt.Errorf("error create before expired")
	}
	return payload.IssuedAt, nil
}

// GetIssuer implements jwt.Claims.
func (payload *Payload) GetIssuer() (string, error) {
	return payload.Issuer, nil
}

// GetNotBefore implements jwt.Claims.
func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return payload.NotBefore, nil
}

// GetSubject implements jwt.Claims.
func (payload *Payload) GetSubject() (string, error) {
	return payload.Subject, nil
}

// NewPayload creates a new token payload with specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssueAt:   time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
