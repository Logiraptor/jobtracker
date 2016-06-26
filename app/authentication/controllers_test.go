package authentication

import (
	"errors"
	"jobtracker/app/doubles"
	"jobtracker/app/models"
	"jobtracker/app/tests"
	"jobtracker/app/web"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/manveru/faker"

	"github.com/stretchr/testify/assert"
)

func TestRegistrationsController(t *testing.T) {
	fake, _ := faker.New("en")
	tests.Describe(t, "RegistrationsController", func(c *tests.Context) {
		var (
			pather      = web.NewPather(&doubles.FakeLogger{}, web.Routes())
			authService = &PasswordAuthService{
				UserRepo:    doubles.NewFakeUserRepository(),
				SessionRepo: doubles.NewFakeSessionRepository(),
				Hasher:      doubles.NewFakePasswordHasher(),
			}
			email, password string
			controller      RegistrationsController
			recorder        *httptest.ResponseRecorder
			request         *http.Request
		)
		c.Before(func() {
			email = fake.Email()
			password = fake.Characters(10)
			controller = RegistrationsController{
				Pather:      pather,
				AuthService: authService,
			}
			recorder = httptest.NewRecorder()
			request = doubles.NewRequest(t, "POST", "/create", url.Values{
				"email":            {email},
				"password":         {password},
				"password_confirm": {password},
			})
		})

		c.Describe("Successful Create", func(c *tests.Context) {
			c.It("Redirects to the index path", func() {
				controller.Create(recorder, request)

				assert.Equal(t, http.StatusFound, recorder.Code)
				assert.Equal(t, web.NewPather(nil, web.Routes()).Path("index"), recorder.HeaderMap.Get("Location"))
			})

			c.It("Sets a session cookie", func() {
				controller.Create(recorder, request)

				assert.NotEmpty(t, recorder.HeaderMap.Get("Set-Cookie"))
			})
		})

		c.Describe("Failed Create", func(c *tests.Context) {
			c.Before(func() {
				fake := doubles.NewFakeUserRepository()
				fake.Store_ = func(models.User) error {
					return errors.New("cannot store user")
				}
				authService.UserRepo = fake
			})

			c.It("Does not set a session cookie", func() {
				controller.Create(recorder, request)

				assert.Empty(t, recorder.HeaderMap.Get("Set-Cookie"))
			})

			c.It("Redirects to the index path", func() {
				controller.Create(recorder, request)

				assert.Equal(t, http.StatusFound, recorder.Code)
				assert.Equal(t, web.NewPather(nil, web.Routes()).Path("index"), recorder.HeaderMap.Get("Location"))
			})
		})
	})
}
