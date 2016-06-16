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
	r.AuthService.Create(models.User{
		Email: req.FormValue("email"),
	}, req.FormValue("password"))
	http.Redirect(rw, req, r.Path("index"), http.StatusFound)
}
