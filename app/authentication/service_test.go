package authentication

import (
	"jobtracker/app/models"
	"jobtracker/app/tests/doubles"

	"jobtracker/app/tests"

	"github.com/manveru/faker"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestPasswordAuthService(t *testing.T) {
	tests.Describe(t, "PasswordAuthService", func(c *tests.Context) {
		var (
			userRepo        *doubles.FakeUserRepository
			hasher          *doubles.FakePasswordHasher
			authService     *PasswordAuthService
			fake, _         = faker.New("en")
			email, password string
		)
		c.Before(func() {
			userRepo = doubles.NewFakeUserRepository()
			hasher = doubles.NewFakePasswordHasher()
			authService = &PasswordAuthService{
				Hasher:   hasher,
				UserRepo: userRepo,
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
				user, err := authService.Authenticate(email, password)
				assert.NoError(t, err)
				assert.Equal(t, email, user.Email)
			})

			c.Describe("Failed Lookup", func(c *tests.Context) {
				c.Before(func() {
					userRepo.FindByEmail_ = func(string) (*models.User, error) {
						return nil, assert.AnError
					}
				})

				c.It("Returns invalid credentials", func() {
					user, err := authService.Authenticate(email, password)
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
					user, err := authService.Authenticate(email, password)
					assert.Nil(t, user)
					assert.Equal(t, ErrInvalidCredentials, err)
				})
			})
		})
	})
}
