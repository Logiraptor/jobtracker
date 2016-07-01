package contracts

import (
	"jobtracker/app/models"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/manveru/faker"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Store(models.User) error
}

func UserRepo(t *testing.T, repo UserRepository) {
	fake, _ := faker.New("en")
	for i := 0; i < 10; i++ {
		email := fake.Email()
		hash := fake.Characters(20)
		token := fake.Characters(20)
		err := repo.Store(models.User{
			Email:        email,
			PasswordHash: hash,
			CurrentToken: token,
		})

		assert.NoError(t, err)

		user, err := repo.FindByEmail(email)
		assert.NoError(t, err)
		assert.Equal(t, user.Email, email)
		assert.Equal(t, user.PasswordHash, hash)
		assert.Equal(t, user.CurrentToken, token)
	}
}

type SessionRepository interface {
	FindByToken(token string) (*models.User, error)
	DeleteByToken(token string) error
	New(models.User) (string, error)
}

func SessionRepo(t *testing.T, users UserRepository, repo SessionRepository) {
	fake, _ := faker.New("en")
	for i := 0; i < 10; i++ {
		email := fake.Email()
		hash := fake.Characters(20)
		token := fake.Characters(20)
		user := models.User{
			Email:        email,
			PasswordHash: hash,
			CurrentToken: token,
		}

		err := users.Store(user)
		assert.NoError(t, err)

		session, err := repo.New(user)
		assert.NoError(t, err)
		assert.NotEmpty(t, session)

		u, err := repo.FindByToken(session)
		assert.NoError(t, err)
		assert.Equal(t, user, *u)

		err = repo.DeleteByToken(session)
		assert.NoError(t, err)

		u, err = repo.FindByToken(session)
		assert.Error(t, err)
		assert.Nil(t, u)
	}
}
