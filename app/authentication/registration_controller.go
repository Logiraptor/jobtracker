package authentication

import (
	"jobtracker/app/models"
	"net/http"

	"github.com/gorilla/mux"
)

type RegistrationsController struct {
	AuthService        AuthService
	HTTPSessionTracker HTTPSessionTracker
}

func NewRegistrationsController(authService AuthService, sessionTracker HTTPSessionTracker) RegistrationsController {
	return RegistrationsController{
		AuthService:        authService,
		HTTPSessionTracker: sessionTracker,
	}
}

func (r RegistrationsController) Create(rw http.ResponseWriter, req *http.Request) {
	email, password := req.FormValue("email"), req.FormValue("password")
	if err := r.AuthService.Create(models.User{Email: email}, password); err == nil {
		r.HTTPSessionTracker.Login(rw, req, models.User{Email: email})
	}

	http.Redirect(rw, req, "/", http.StatusFound)
}

func (r RegistrationsController) Register(mux *mux.Router) {
	mux.Path("/sign_up").Methods("POST").HandlerFunc(r.Create)
}
