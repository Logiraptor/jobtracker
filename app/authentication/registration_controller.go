package authentication

import (
	"jobtracker/app/models"
	"jobtracker/app/web"
	"net/http"
)

type RegistrationsController struct {
	web.Pather
	AuthService        AuthService
	HTTPSessionTracker HTTPSessionTracker
}

func (r RegistrationsController) Create(rw http.ResponseWriter, req *http.Request) {
	email, password := req.FormValue("email"), req.FormValue("password")
	if err := r.AuthService.Create(models.User{Email: email}, password); err == nil {
		r.HTTPSessionTracker.Login(rw, req, models.User{Email: email})
	}

	http.Redirect(rw, req, r.Path("index"), http.StatusFound)
}
