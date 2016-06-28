package authentication

import (
	"jobtracker/app/models"
	"jobtracker/app/tests"
	"jobtracker/app/tests/doubles"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"testing"
)

type SpyHandler struct {
	called bool
}

func (s *SpyHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.called = true
}

func TestMiddleware(t *testing.T) {
	tests.Describe(t, "Middleware", func(c *tests.Context) {
		c.Describe("RequireAuthentication", func(c *tests.Context) {
			var stubSessionTracker *doubles.StubHTTPSessionTracker
			var fakeAuthedHandler *SpyHandler
			var fakeUnauthedHandler *SpyHandler
			var wrappedHandler http.HandlerFunc
			var recorder *httptest.ResponseRecorder
			c.Before(func() {
				stubSessionTracker = doubles.NewStubHTTPSessionTracker()
				fakeAuthedHandler = new(SpyHandler)
				fakeUnauthedHandler = new(SpyHandler)
				wrappedHandler = RequireAuthentication(stubSessionTracker, fakeAuthedHandler, fakeUnauthedHandler)
				recorder = httptest.NewRecorder()
			})

			c.Describe("Authenticated request", func(c *tests.Context) {
				c.Before(func() {
					stubSessionTracker.CurrentUser_ = func(*http.Request) (*models.User, bool) {
						return new(models.User), true
					}
				})
				c.It("Forwards the request through", func() {
					req := doubles.NewRequest(t, "GET", "/", nil)
					wrappedHandler.ServeHTTP(recorder, req)
					assert.Equal(t, http.StatusOK, recorder.Code)
					assert.True(t, fakeAuthedHandler.called)
					assert.False(t, fakeUnauthedHandler.called)
				})
			})

			c.Describe("Unauthenticated request", func(c *tests.Context) {
				c.Before(func() {
					stubSessionTracker.CurrentUser_ = func(*http.Request) (*models.User, bool) {
						return nil, false
					}
				})

				c.It("Denies the request", func() {
					req := doubles.NewRequest(t, "GET", "/", nil)
					wrappedHandler.ServeHTTP(recorder, req)
					assert.Equal(t, http.StatusOK, recorder.Code)
					assert.True(t, fakeUnauthedHandler.called)
					assert.False(t, fakeAuthedHandler.called)
				})
			})
		})
	})
}
