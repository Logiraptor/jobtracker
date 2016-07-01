package doubles

import (
	"jobtracker/app/models"
	"net/http"
)

type StubHTTPSessionTracker struct {
	Login_       func(http.ResponseWriter, *http.Request, models.User) error
	Logout_      func(http.ResponseWriter, *http.Request) error
	CurrentUser_ func(*http.Request) (*models.User, bool)
}

func NewStubHTTPSessionTracker() *StubHTTPSessionTracker {
	return &StubHTTPSessionTracker{
		Login_: func(http.ResponseWriter, *http.Request, models.User) error {
			return nil
		},
		Logout_: func(http.ResponseWriter, *http.Request) error {
			return nil
		},
		CurrentUser_: func(*http.Request) (*models.User, bool) {
			return nil, false
		},
	}
}

func (f *StubHTTPSessionTracker) Login(rw http.ResponseWriter, req *http.Request, user models.User) error {
	return f.Login_(rw, req, user)
}

func (f *StubHTTPSessionTracker) Logout(rw http.ResponseWriter, req *http.Request) error {
	return f.Logout_(rw, req)
}

func (f *StubHTTPSessionTracker) CurrentUser(req *http.Request) (*models.User, bool) {
	return f.CurrentUser_(req)
}
