package authentication

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type SessionsController struct {
	Logger         *logrus.Logger
	AuthService    AuthService
	SessionTracker HTTPSessionTracker
}

func NewSessionsController(logger *logrus.Logger, authService AuthService, sessionTracker HTTPSessionTracker) SessionsController {
	return SessionsController{
		Logger:         logger,
		AuthService:    authService,
		SessionTracker: sessionTracker,
	}
}

func (s SessionsController) Create(rw http.ResponseWriter, req *http.Request) {
	email, password := req.FormValue("email"), req.FormValue("password")
	user, err := s.AuthService.Authenticate(email, password)
	if err == nil {
		err := s.SessionTracker.Login(rw, req, *user)
		if err != nil {
			s.Logger.WithError(err).Print("Login failed")
		}
		s.Logger.WithFields(logrus.Fields{
			"email":  email,
			"cookie": rw.Header().Get("Set-Cookie"),
		}).Print("Login")
	} else {
		s.Logger.WithError(err).Print("Authentication")
	}
	http.Redirect(rw, req, "/", http.StatusFound)
}

func (s SessionsController) Destroy(rw http.ResponseWriter, req *http.Request) {
	s.SessionTracker.Logout(rw, req)
	http.Redirect(rw, req, "/", http.StatusFound)
}

func (s SessionsController) Register(mux *mux.Router) {
	mux.Path("/sign_in").Methods("POST").HandlerFunc(s.Create)
	mux.Path("/sign_out").HandlerFunc(s.Destroy)
}
