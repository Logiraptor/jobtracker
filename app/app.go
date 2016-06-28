package app

import (
	"html/template"
	"jobtracker/app/authentication"
	"jobtracker/app/inject"
	"jobtracker/app/web"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/securecookie"

	"github.com/gorilla/sessions"
)

type Context struct {
	Logger  web.Logger
	AppRoot string
	Port    int
}

func Start(ctx Context) error {

	var routes = web.Routes()

	cookieStore := sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

	c := inject.NewContainer()
	c.Register((*sessions.Store)(nil), cookieStore)
	c.Register((*web.Logger)(nil), ctx.Logger)
	c.Register((*web.Pather)(nil), web.NewPather(ctx.Logger, routes))
	c.Register((*authentication.HTTPSessionTracker)(nil), authentication.CookieSessionTracker{
		SessionName: "jobtracker",
	})
	c.Register((*authentication.SessionRepository)(nil), TODO)
	c.Register((*authentication.UserRepository)(nil), TODO)
	c.Register((*authentication.PasswordHasher)(nil), authentication.BCryptPasswordHasher{})
	c.Register((*authentication.AuthService)(nil), authentication.PasswordAuthService{})

	var pdfController PdfController
	if err := c.FillStruct(&pdfController); err != nil {
		return err
	}

	var registrationsController authentication.RegistrationsController
	if err := c.FillStruct(&registrationsController); err != nil {
		return err
	}

	var tmpls, err = template.ParseGlob(filepath.Join(ctx.AppRoot, "public/*.html"))
	if err != nil {
		return err
	}

	var templateController = TemplateController{
		AppContext: ctx,
		Template:   tmpls,
	}

	routes.Get("generate_pdf").HandlerFunc(pdfController.Generate)
	routes.Get("sign_up").HandlerFunc(registrationsController.Create)

	routes.Get("index").Handler(templateController)

	ctx.Logger.Log("App started on port: %d", ctx.Port)

	return http.ListenAndServe(":"+strconv.Itoa(ctx.Port), routes)
}
