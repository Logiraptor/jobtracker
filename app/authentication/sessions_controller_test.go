package authentication

import (
	"io/ioutil"
	"jobtracker/app/models"
	"jobtracker/app/tests"
	"jobtracker/app/tests/doubles"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/manveru/faker"

	"github.com/stretchr/testify/assert"
)

func TestSessionsController(t *testing.T) {
	tests.Describe(t, "SessionsController", func(c *tests.Context) {
		var (
			user       models.User
			request    *http.Request
			recorder   *httptest.ResponseRecorder
			controller SessionsController
			hasher     *doubles.FakePasswordHasher
			userRepo   *doubles.FakeUserRepository
			sessRepo   *doubles.FakeSessionRepository
			logger     = logrus.New()
		)
		logger.Out = ioutil.Discard
		c.Before(func() {
			fake, _ := faker.New("en")
			user = models.User{
				Email:        fake.Email(),
				PasswordHash: fake.Characters(20),
			}

			recorder = httptest.NewRecorder()
			hasher = doubles.NewFakePasswordHasher()
			userRepo = doubles.NewFakeUserRepository()
			sessRepo = doubles.NewFakeSessionRepository()

			var sessionStore = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))
			var sessionTracker = NewCookieSessionTracker("jobtracker", logger, sessionStore, sessRepo)
			var authService = NewPasswordAuthService(hasher, userRepo)
			controller = NewSessionsController(logger, authService, sessionTracker)
		})

		c.Describe("#Create", func(c *tests.Context) {
			c.Before(func() {
				request = doubles.NewRequest(t, "POST", "/login", url.Values{
					"email":    {"email@example.com"},
					"password": {"password"},
				})
			})

			c.Describe("Successful login", func(c *tests.Context) {
				c.Before(func() {
					userRepo.FindByEmail_ = func(string) (*models.User, error) {
						return &user, nil
					}
					hasher.Verify_ = func(_, _ string) bool {
						return true
					}
				})
				c.It("Sets a session cookie", func() {
					controller.Create(recorder, request)
					assert.NotEmpty(t, recorder.HeaderMap.Get("Set-Cookie"))
				})
				c.It("Redirects to the index", func() {
					controller.Create(recorder, request)
					assert.Equal(t, http.StatusFound, recorder.Code)
					assert.Equal(t, "/", recorder.HeaderMap.Get("Location"))
				})
			})

			c.Describe("Invalid password", func(c *tests.Context) {
				c.Before(func() {
					userRepo.FindByEmail_ = func(string) (*models.User, error) {
						return nil, assert.AnError
					}
					hasher.Verify_ = func(_, _ string) bool {
						return false
					}
				})

				c.It("Does not set a session cookie", func() {
					controller.Create(recorder, request)
					assert.Empty(t, recorder.HeaderMap.Get("Set-Cookie"))
				})
				c.It("Redirects to the index", func() {
					controller.Create(recorder, request)
					assert.Equal(t, http.StatusFound, recorder.Code)
					assert.Equal(t, "/", recorder.HeaderMap.Get("Location"))
				})
			})
		})

		c.Describe("#Destroy", func(c *tests.Context) {
			c.It("Redirects to the index", func() {
				controller.Destroy(recorder, request)
				assert.Equal(t, http.StatusFound, recorder.Code)
				assert.Equal(t, "/", recorder.HeaderMap.Get("Location"))
			})
		})
	})
}
