package app

import (
	"jobtracker/app/models"
	"jobtracker/app/services"
	"net/http"
)

type RegistrationsController struct {
	Pather
	services.AuthService
}

func (r RegistrationsController) Create(rw http.ResponseWriter, req *http.Request) {
	email, password := req.FormValue("email"), req.FormValue("password")
	if err := r.AuthService.Create(models.User{Email: email}, password); err == nil {
		r.AuthService.Authenticate(email, password)
		rw.Header().Add("Set-Cookie", "session")
	}

	http.Redirect(rw, req, r.Path("index"), http.StatusFound)
}
