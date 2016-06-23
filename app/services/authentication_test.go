package services

import (
	"errors"
	"jobtracker/app/mocks"
	"jobtracker/app/models"
	"jobtracker/app/tests"

	"github.com/golang/mock/gomock"
	"github.com/manveru/faker"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestPasswordAuthServiceCreate(t *testing.T) {
	var (
		ctrl        = gomock.NewController(t)
		repo        = mocks.NewMockUserRepository(ctrl)
		hasher      = mocks.NewMockPasswordHasher(ctrl)
		authService = &PasswordAuthService{
			Hasher:   hasher,
			UserRepo: repo,
		}
		fake, _        = faker.New("en")
		email          = fake.Email()
		password       = fake.Characters(20)
		hashedPassword = fake.Characters(20)
		errorMessage   = fake.Characters(20)
	)
	defer ctrl.Finish()

	hasher.EXPECT().New(password).Return(hashedPassword)

	repo.EXPECT().Store(models.User{
		Email:        email,
		PasswordHash: hashedPassword,
	}).Return(errors.New(errorMessage))

	err := authService.Create(models.User{
		Email: email,
	}, password)
	assert.NotNil(t, err)
	assert.Equal(t, errorMessage, err.Error())
}

func TestPasswordAuthServiceAuthenticate(t *testing.T) {
	var (
		ctrl                            = gomock.NewController(t)
		repo                            *mocks.MockUserRepository
		hasher                          *mocks.MockPasswordHasher
		authService                     *PasswordAuthService
		fake, _                         = faker.New("en")
		email, password, storedPassword string
	)
	defer ctrl.Finish()
	tests.Describe(t, "PasswordAuthService#Authentication", func(c *tests.Context) {
		c.Before(func() {
			ctrl = gomock.NewController(t)
			repo = mocks.NewMockUserRepository(ctrl)
			hasher = mocks.NewMockPasswordHasher(ctrl)
			authService = &PasswordAuthService{Hasher: hasher, UserRepo: repo}
			email = fake.Email()
			password = fake.Characters(20)
			storedPassword = fake.Characters(20)
		})

		c.Describe("Successful lookup", func(c *tests.Context) {
			c.Before(func() {
				repo.EXPECT().FindByEmail(email).Return(&models.User{
					Email:        email,
					PasswordHash: storedPassword,
				}, nil)
				hasher.EXPECT().Verify(storedPassword, password).Return(true)
			})

			c.It("Returns the associated user", func() {
				user, err := authService.Authenticate(email, password)
				assert.Nil(t, err)
				assert.EqualValues(t, &models.User{
					Email:        email,
					PasswordHash: storedPassword,
				}, user)
			})
		})

		c.Describe("Failed lookup", func(c *tests.Context) {
			var errorMessage string
			c.Before(func() {
				errorMessage = fake.Characters(20)
				repo.EXPECT().FindByEmail(email).Return(nil, errors.New(errorMessage))
			})

			c.It("Returns the error message", func() {
				user, err := authService.Authenticate(email, password)
				assert.Nil(t, user)
				assert.NotNil(t, err)
				assert.Equal(t, ErrInvalidCredentials, err)
			})
		})

		c.Describe("Invalid Password", func(c *tests.Context) {
			c.Before(func() {
				repo.EXPECT().FindByEmail(email).Return(&models.User{
					Email:        email,
					PasswordHash: storedPassword,
				}, nil)
				hasher.EXPECT().Verify(storedPassword, password).Return(false)
			})

			c.It("Returns the error message", func() {
				user, err := authService.Authenticate(email, password)
				assert.Nil(t, user)
				assert.NotNil(t, err)
				assert.Equal(t, ErrInvalidCredentials, err)
			})
		})
	})
}
