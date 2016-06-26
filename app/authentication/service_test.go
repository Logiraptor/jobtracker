package authentication

import (
	"jobtracker/app/doubles"
	"jobtracker/app/models"

	"jobtracker/app/tests"

	"github.com/manveru/faker"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestPasswordAuthService(t *testing.T) {
	tests.Describe(t, "PasswordAuthService", func(c *tests.Context) {
		var (
			userRepo        *doubles.FakeUserRepository
			sessionRepo     *doubles.FakeSessionRepository
			hasher          *doubles.FakePasswordHasher
			authService     *PasswordAuthService
			fake, _         = faker.New("en")
			email, password string
		)
		c.Before(func() {
			userRepo = doubles.NewFakeUserRepository()
			sessionRepo = doubles.NewFakeSessionRepository()
			hasher = doubles.NewFakePasswordHasher()
			authService = &PasswordAuthService{
				Hasher:      hasher,
				UserRepo:    userRepo,
				SessionRepo: sessionRepo,
			}
			email = fake.Email()
			password = fake.Characters(20)
		})

		c.Describe("#Create", func(c *tests.Context) {
			c.Before(func() {
				_, err := userRepo.FindByEmail(email)
				assert.Error(t, err)
			})

			c.It("Creates a user in the repo with a password hash", func() {
				err := authService.Create(models.User{
					Email: email,
				}, password)
				assert.NoError(t, err)

				user, err := userRepo.FindByEmail(email)
				assert.NoError(t, err)
				assert.Equal(t, email, user.Email)
				assert.NotEqual(t, password, user.PasswordHash)
				assert.NotEmpty(t, user.PasswordHash)
			})
		})
		c.Describe("#Authenticate", func(c *tests.Context) {
			c.Before(func() {
				authService.Create(models.User{Email: email}, password)
			})

			c.It("Returns the created user and a token", func() {
				user, token, err := authService.Authenticate(email, password)
				assert.NoError(t, err)
				assert.Equal(t, email, user.Email)
				assert.NotEmpty(t, token)
				tokenUser, err := sessionRepo.FindByToken(token)
				assert.NoError(t, err)
				assert.Equal(t, email, tokenUser.Email)
			})

			c.Describe("Failed Lookup", func(c *tests.Context) {
				c.Before(func() {
					userRepo.FindByEmail_ = func(string) (*models.User, error) {
						return nil, assert.AnError
					}
				})

				c.It("Returns invalid credentials", func() {
					user, token, err := authService.Authenticate(email, password)
					assert.Empty(t, token)
					assert.Nil(t, user)
					assert.Equal(t, ErrInvalidCredentials, err)
				})
			})

			c.Describe("Invalid Password", func(c *tests.Context) {
				c.Before(func() {
					hasher.Verify_ = func(string, string) bool {
						return false
					}
				})

				c.It("Returns invalid credentials", func() {
					user, token, err := authService.Authenticate(email, password)
					assert.Empty(t, token)
					assert.Nil(t, user)
					assert.Equal(t, ErrInvalidCredentials, err)
				})
			})

			c.Describe("Failed session creation", func(c *tests.Context) {
				c.Before(func() {
					sessionRepo.New_ = func(models.User) (string, error) {
						return "", assert.AnError
					}
				})

				c.It("Returns the error from the session repo", func() {
					user, token, err := authService.Authenticate(email, password)
					assert.Empty(t, token)
					assert.Nil(t, user)
					assert.Equal(t, assert.AnError, err)
				})
			})
		})
	})
}
