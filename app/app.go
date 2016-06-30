package app

import (
	"html/template"
	"jobtracker/app/authentication"
	"jobtracker/app/web"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

	_ "github.com/lib/pq"

	"database/sql"
)

type Context struct {
	Logger  web.Logger
	AppRoot string
	Port    int
}

type Controller interface {
	Register(mux *mux.Router)
}

func Start(ctx Context) error {
	db, err := sql.Open("postgres", "postgres://jobtracker:@localhost:5433/jobtracker")
	if err != nil {
		return err
	}

	var (
		store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))
		tmpls = template.Must(template.ParseGlob(filepath.Join(ctx.AppRoot, "public/*.html")))

		userRepo                = authentication.NewPSQLUserRepo(db)
		sessionRepository       = authentication.NewPSQLSessionRepo(db)
		hasher                  = authentication.NewBCryptPasswordHasher(10)
		httpSessionTracker      = authentication.NewCookieSessionTracker("jobtracker", store, sessionRepository)
		authService             = authentication.NewPasswordAuthService(hasher, userRepo)
		registrationsController = authentication.NewRegistrationsController(authService, httpSessionTracker)

		pdfController       = NewPdfController(ctx.Logger)
		dashboardController = NewDashboardController(tmpls, authService)
	)

	var controllers = []Controller{
		registrationsController,
		pdfController,
		dashboardController,
	}

	router := mux.NewRouter()
	for _, c := range controllers {
		c.Register(router)
	}

	ctx.Logger.Log("App started on port: %d", ctx.Port)
	return http.ListenAndServe(":"+strconv.Itoa(ctx.Port), router)
}
