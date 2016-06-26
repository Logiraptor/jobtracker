package app

import (
	"jobtracker/app/mocks"
	"jobtracker/app/models"
	"jobtracker/app/services"
	"jobtracker/app/tests"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/manveru/faker"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestRegistrationsController(t *testing.T) {
	fake, _ := faker.New("en")
	tests.Describe(t, "RegistrationsController", func(c *tests.Context) {
		var (
			pather                        *mocks.MockPather
			authService                   *mocks.MockAuthService
			email, password, sessionToken string
			controller                    RegistrationsController
			recorder                      *httptest.ResponseRecorder
			request                       *http.Request
			mockCtrl                      *gomock.Controller
		)
		c.Before(func() {
			mockCtrl = gomock.NewController(t)
			pather = mocks.NewMockPather(mockCtrl)
			pather.EXPECT().Path("index").Return("fake path")

			email = fake.Email()
			password = fake.Characters(10)
			authService = mocks.NewMockAuthService(mockCtrl)
			controller = RegistrationsController{
				Pather:      pather,
				AuthService: authService,
			}
			recorder = httptest.NewRecorder()
			request = mustNewRequest(t, "POST", "/create", url.Values{
				"email":            {email},
				"password":         {password},
				"password_confirm": {password},
			})
		})

		c.After(func() {
			mockCtrl.Finish()
		})

		c.Describe("Successful Create", func(c *tests.Context) {
			c.Before(func() {
				authService.EXPECT().Create(models.User{
					Email: email,
				}, password).Return(nil)
				authService.EXPECT().Authenticate(email, password).Return(&models.User{
					Email:        email,
					PasswordHash: password,
				}, sessionToken, nil)
			})

			c.It("Redirects to the index path", func() {
				controller.Create(recorder, request)

				assert.Equal(t, http.StatusFound, recorder.Code)
				assert.Equal(t, "/fake path", recorder.HeaderMap.Get("Location"))
			})

			c.It("Sets a session cookie", func() {
				controller.Create(recorder, request)

				assert.NotEmpty(t, recorder.HeaderMap.Get("Set-Cookie"))
			})
		})

		c.Describe("Failed Create", func(c *tests.Context) {
			c.Before(func() {
				authService.EXPECT().Create(models.User{
					Email: email,
				}, password).Return(services.ErrInvalidCredentials)
			})

			c.It("Redirects to the index path", func() {
				controller.Create(recorder, request)

				assert.Equal(t, http.StatusFound, recorder.Code)
				assert.Equal(t, "/fake path", recorder.HeaderMap.Get("Location"))
			})
		})
	})
}
