package authentication

import (
	"net/http"

	"github.com/gorilla/mux"
)

type SessionsController struct {
	AuthService    AuthService
	SessionTracker HTTPSessionTracker
}

func NewSessionsController(authService AuthService, sessionTracker HTTPSessionTracker) SessionsController {
	return SessionsController{
		AuthService:    authService,
		SessionTracker: sessionTracker,
	}
}

func (s SessionsController) Create(rw http.ResponseWriter, req *http.Request) {
	email, password := req.FormValue("email"), req.FormValue("password")
	user, err := s.AuthService.Authenticate(email, password)
	if err == nil {
		s.SessionTracker.Login(rw, req, *user)
	}
	http.Redirect(rw, req, "/", http.StatusFound)
}

func (s SessionsController) Register(mux *mux.Router) {
	mux.Path("/login").Methods("POST").HandlerFunc(s.Create)
}
