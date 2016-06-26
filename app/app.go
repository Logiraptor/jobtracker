package app

import (
	"html/template"
	"jobtracker/app/authentication"
	"jobtracker/app/web"
	"net/http"
	"path/filepath"
	"strconv"
)

type Context struct {
	Logger  web.Logger
	AppRoot string
	Port    int
}

func Start(ctx Context) error {
	var pdfController = PdfController{
		Logger: ctx.Logger,
	}

	var routes = web.Routes()

	var registrationsController = authentication.RegistrationsController{
		Pather:      web.NewPather(ctx.Logger, routes),
		AuthService: (authentication.AuthService)(nil),
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
