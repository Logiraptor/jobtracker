package services

import (
	"errors"
	"jobtracker/app/models"
)

var (
	ErrInvalidCredentials = errors.New("invalid email and/or password")
)

//go:generate mockgen -source $GOFILE -destination ../mocks/mock_services.go -package mocks

type AuthService interface {
	Create(user models.User, password string) error
	Authenticate(email, password string) (*models.User, error)
}

type PasswordHasher interface {
	New(password string) string
	Verify(hash, password string) bool
}

type PasswordAuthService struct {
	Hasher   PasswordHasher
	UserRepo models.UserRepository
}

func (b PasswordAuthService) Create(user models.User, password string) error {
	user.PasswordHash = b.Hasher.New(password)
	return b.UserRepo.Store(user)
}

func (b PasswordAuthService) Authenticate(email, password string) (*models.User, error) {
	user, err := b.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	if b.Hasher.Verify(user.PasswordHash, password) {
		return user, nil
	}
	return nil, ErrInvalidCredentials
}
