package app

import (
	"html/template"
	"jobtracker/app/authentication"
	"jobtracker/app/web"
	"net/http"
	"path/filepath"
	"strconv"

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

func Start(ctx Context) error {

	var routes = web.Routes()

	db, err := sql.Open("postgres", "postgres://jobtracker:@localhost:5433/jobtracker")
	if err != nil {
		return err
	}

	var (
		pather = web.NewPather(ctx.Logger, routes)
		store  = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))
		tmpls  = template.Must(template.ParseGlob(filepath.Join(ctx.AppRoot, "public/*.html")))

		pdfController      = NewPdfController(ctx.Logger)
		templateController = NewTemplateController(tmpls, ctx)

		userRepo                = authentication.NewPSQLUserRepo(db)
		sessionRepository       = authentication.NewPSQLSessionRepo(db)
		hasher                  = authentication.NewBCryptPasswordHasher(10)
		httpSessionTracker      = authentication.NewCookieSessionTracker("jobtracker", store, sessionRepository)
		authService             = authentication.NewPasswordAuthService(hasher, userRepo)
		registrationsController = authentication.NewRegistrationsController(pather, authService, httpSessionTracker)
	)

	routes.Get("generate_pdf").HandlerFunc(pdfController.Generate)
	routes.Get("sign_up").HandlerFunc(registrationsController.Create)

	routes.Get("index").Handler(templateController)

	ctx.Logger.Log("App started on port: %d", ctx.Port)

	return http.ListenAndServe(":"+strconv.Itoa(ctx.Port), routes)
}
