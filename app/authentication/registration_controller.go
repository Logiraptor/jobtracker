package authentication

import (
	"jobtracker/app/models"
	"jobtracker/app/web"
	"net/http"

	"github.com/gorilla/mux"
)

type RegistrationsController struct {
	View               web.View
	AuthService        AuthService
	HTTPSessionTracker HTTPSessionTracker
}

func NewRegistrationsController(view web.View, authService AuthService, sessionTracker HTTPSessionTracker) RegistrationsController {
	return RegistrationsController{
		View:               view,
		AuthService:        authService,
		HTTPSessionTracker: sessionTracker,
	}
}

func (r RegistrationsController) New(rw http.ResponseWriter, req *http.Request) {
	r.View.Render(rw, "sign_up.html", nil)
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
	mux.Path("/sign_up").Methods("GET").HandlerFunc(r.New)
}
