package app

import (
	"encoding/hex"
	"html/template"
	"jobtracker/app/authentication"
	"jobtracker/app/web"
	"net/http"
	"path/filepath"
	"strconv"

	log "github.com/Sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	_ "github.com/lib/pq"

	"database/sql"
)

type Context struct {
	Logger  *log.Logger
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
		authKey, _ = hex.DecodeString("3e408db5b476dcda920383089e765df923e1a7b845ecb30098d41b56228eba3cc81e823200b3ea50d0682ebc8b604ddba72189392b312f7c292b08360349558b")
		encKey, _  = hex.DecodeString("09d04f2aa1d2bf8103a3b8dc736c76c1a3fe55d9187c25f868cdb486784dabd0")
		store      = sessions.NewCookieStore(authKey, encKey)
		tmpls      = template.Must(template.ParseGlob(filepath.Join(ctx.AppRoot, "public/*.html")))
		view       = web.NewTemplateView(tmpls)

		userRepo                = authentication.NewPSQLUserRepo(db)
		sessionRepository       = authentication.NewPSQLSessionRepo(db)
		hasher                  = authentication.NewBCryptPasswordHasher(10)
		httpSessionTracker      = authentication.NewCookieSessionTracker("jobtracker", ctx.Logger, store, sessionRepository)
		authService             = authentication.NewPasswordAuthService(hasher, userRepo)
		registrationsController = authentication.NewRegistrationsController(view, authService, httpSessionTracker)
		sessionsController      = authentication.NewSessionsController(ctx.Logger, authService, httpSessionTracker)

		pdfController       = NewPdfController(ctx.Logger)
		dashboardController = NewDashboardController(view, httpSessionTracker)
	)

	var controllers = []Controller{
		pdfController,
		registrationsController,
		sessionsController,
		dashboardController,
	}

	router := mux.NewRouter()
	for _, c := range controllers {
		c.Register(router)
	}

	ctx.Logger.WithField("port", ctx.Port).Print("App started")
	return http.ListenAndServe(":"+strconv.Itoa(ctx.Port), web.LoggerMiddleware(ctx.Logger, router))
}
