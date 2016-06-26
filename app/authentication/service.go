package authentication

import (
	"errors"
	"jobtracker/app/models"
)

var (
	ErrInvalidCredentials = errors.New("invalid email and/or password")
)

type AuthService interface {
	Create(user models.User, password string) error
	Authenticate(email, password string) (user *models.User, token string, err error)
}

type PasswordHasher interface {
	New(password string) string
	Verify(hash, password string) bool
}

type PasswordAuthService struct {
	Hasher      PasswordHasher
	UserRepo    UserRepository
	SessionRepo SessionRepository
}

func (b PasswordAuthService) Create(user models.User, password string) error {
	user.PasswordHash = b.Hasher.New(password)
	return b.UserRepo.Store(user)
}

func (b PasswordAuthService) Authenticate(email, password string) (*models.User, string, error) {
	user, err := b.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}
	if !b.Hasher.Verify(user.PasswordHash, password) {
		return nil, "", ErrInvalidCredentials
	}
	token, err := b.SessionRepo.New(*user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
