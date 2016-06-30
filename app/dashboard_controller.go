package app

import (
	"jobtracker/app/authentication"
	"jobtracker/app/web"
	"net/http"

	"github.com/gorilla/mux"
)

type DashboardController struct {
	View               web.View
	HTTPSessionTracker authentication.HTTPSessionTracker
}

func NewDashboardController(view web.View, httpSessionTracker authentication.HTTPSessionTracker) DashboardController {
	return DashboardController{
		View:               view,
		HTTPSessionTracker: httpSessionTracker,
	}
}

func (d DashboardController) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	user, _ := d.HTTPSessionTracker.CurrentUser(req)
	d.View.Render(rw, "index.html", user)
}

func (d DashboardController) Register(mux *mux.Router) {
	mux.Path("/").Handler(d)
}
