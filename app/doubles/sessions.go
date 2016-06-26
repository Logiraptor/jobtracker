package doubles

import (
	"encoding/hex"
	"errors"
	"jobtracker/app/models"
	"jobtracker/app/services"
	"math/rand"
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
			var buf [16]byte
			rand.Read(buf[:])
			token := hex.EncodeToString(buf[:])
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

func NewFakePasswordHasher() *FakePasswordHasher {
	return &FakePasswordHasher{
		New_: func(password string) string {
			return password + "hash"
		},
		Verify_: func(hash, password string) bool {
			return password+"hash" == hash
		},
	}
}

func (f *FakePasswordHasher) New(password string) string {
	return f.New_(password)
}

func (f *FakePasswordHasher) Verify(hash, password string) bool {
	return f.Verify_(hash, password)
}

var _ models.UserRepository = &FakeUserRepository{}
var _ models.SessionRepository = &FakeSessionRepository{}
var _ services.PasswordHasher = &FakePasswordHasher{}
