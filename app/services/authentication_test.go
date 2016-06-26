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
		ctrl                                          = gomock.NewController(t)
		userRepo                                      *mocks.MockUserRepository
		sessionRepo                                   *mocks.MockSessionRepository
		hasher                                        *mocks.MockPasswordHasher
		authService                                   *PasswordAuthService
		fake, _                                       = faker.New("en")
		email, password, storedPassword, sessionToken string
	)
	defer ctrl.Finish()
	tests.Describe(t, "PasswordAuthService#Authentication", func(c *tests.Context) {
		c.Before(func() {
			ctrl = gomock.NewController(t)
			userRepo = mocks.NewMockUserRepository(ctrl)
			sessionRepo = mocks.NewMockSessionRepository(ctrl)
			hasher = mocks.NewMockPasswordHasher(ctrl)
			authService = &PasswordAuthService{Hasher: hasher, UserRepo: userRepo, SessionRepo: sessionRepo}
			email = fake.Email()
			password = fake.Characters(20)
			storedPassword = fake.Characters(20)
			sessionToken = fake.Characters(20)
		})

		c.Describe("Successful lookup", func(c *tests.Context) {
			c.Before(func() {
				userRepo.EXPECT().FindByEmail(email).Return(&models.User{
					Email:        email,
					PasswordHash: storedPassword,
				}, nil)
				sessionRepo.EXPECT().New(models.User{
					Email:        email,
					PasswordHash: storedPassword,
				}).Return(sessionToken, nil)
				hasher.EXPECT().Verify(storedPassword, password).Return(true)
			})

			c.It("Returns the associated user and token", func() {
				user, token, err := authService.Authenticate(email, password)
				assert.Nil(t, err)
				assert.EqualValues(t, &models.User{
					Email:        email,
					PasswordHash: storedPassword,
				}, user)
				assert.Equal(t, sessionToken, token)
			})
		})

		c.Describe("Failed lookup", func(c *tests.Context) {
			var errorMessage string
			c.Before(func() {
				errorMessage = fake.Characters(20)
				userRepo.EXPECT().FindByEmail(email).Return(nil, errors.New(errorMessage))
			})

			c.It("Returns the error message", func() {
				user, token, err := authService.Authenticate(email, password)
				assert.Nil(t, user)
				assert.Empty(t, token)
				assert.NotNil(t, err)
				assert.Equal(t, ErrInvalidCredentials, err)
			})
		})

		c.Describe("Invalid Password", func(c *tests.Context) {
			c.Before(func() {
				userRepo.EXPECT().FindByEmail(email).Return(&models.User{
					Email:        email,
					PasswordHash: storedPassword,
				}, nil)
				hasher.EXPECT().Verify(storedPassword, password).Return(false)
			})

			c.It("Returns the error message", func() {
				user, token, err := authService.Authenticate(email, password)
				assert.Empty(t, token)
				assert.Nil(t, user)
				assert.NotNil(t, err)
				assert.Equal(t, ErrInvalidCredentials, err)
			})
		})

		c.Describe("Failed Session Creation", func(c *tests.Context) {
			var errorMessage = fake.Characters(10)
			c.Before(func() {
				userRepo.EXPECT().FindByEmail(gomock.Any()).Return(&models.User{
					Email:        email,
					PasswordHash: storedPassword,
				}, nil).AnyTimes()
				hasher.EXPECT().Verify(gomock.Any(), gomock.Any()).Return(true).AnyTimes()

				sessionRepo.EXPECT().New(models.User{
					Email:        email,
					PasswordHash: storedPassword,
				}).Return("", errors.New(errorMessage))
			})

			c.It("Returns the session error", func() {
				user, token, err := authService.Authenticate(email, password)
				assert.Empty(t, token)
				assert.Nil(t, user)
				assert.NotNil(t, err)
				assert.Equal(t, errorMessage, err.Error())
			})
		})
	})
}
