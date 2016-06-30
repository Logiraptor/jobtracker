package app

import (
	"html/template"
	"jobtracker/app/authentication"
	"net/http"

	"github.com/gorilla/mux"
)

type DashboardController struct {
	template    *template.Template
	authService authentication.AuthService
}

func NewDashboardController(tmpls *template.Template, authService authentication.AuthService) DashboardController {
	return DashboardController{
		template:    tmpls,
		authService: authService,
	}
}

func (d DashboardController) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

}

func (d DashboardController) Register(mux *mux.Router) {
	mux.Path("/").Handler(d)
}
