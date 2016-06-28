package doubles

import (
	"errors"
	"jobtracker/app/models"

	"strconv"
)

type FakeUserRepository struct {
	FindByEmail_ func(string) (*models.User, error)
	Store_       func(models.User) error
}

func NewFakeUserRepository() *FakeUserRepository {
	var users = make(map[string]models.User)
	return &FakeUserRepository{
		FindByEmail_: func(email string) (*models.User, error) {
			user, ok := users[email]
			if !ok {
				return nil, errors.New("No such user")
			}
			return &user, nil
		},
		Store_: func(user models.User) error {
			users[user.Email] = user
			return nil
		},
	}
}

func (f *FakeUserRepository) FindByEmail(email string) (*models.User, error) {
	return f.FindByEmail_(email)
}

func (f *FakeUserRepository) Store(user models.User) error {
	return f.Store_(user)
}

type FakeSessionRepository struct {
	FindByToken_ func(token string) (*models.User, error)
	New_         func(user models.User) (string, error)
}

func NewFakeSessionRepository() *FakeSessionRepository {
	var sessions = make(map[string]models.User)
	return &FakeSessionRepository{
		FindByToken_: func(token string) (*models.User, error) {
			user, ok := sessions[token]
			if !ok {
				return nil, errors.New("No such user")
			}
			return &user, nil
		},
		New_: func(user models.User) (string, error) {
			token := strconv.Itoa(len(sessions))
			sessions[token] = user
			return token, nil
		},
	}
}

func (f *FakeSessionRepository) FindByToken(token string) (*models.User, error) {
	return f.FindByToken_(token)
}

func (f *FakeSessionRepository) New(user models.User) (string, error) {
	return f.New_(user)
}

type FakePasswordHasher struct {
	New_    func(string) string
	Verify_ func(string, string) bool
}
