package app

import (
	"html/template"
	"jobtracker/app/services"
	"net/http"
	"path/filepath"
	"strconv"
)

//go:generate mockgen -source $GOFILE -destination ./mocks/mock_app.go -package mocks

type Logger interface {
	Log(format string, args ...interface{})
}

type Pather interface {
	Path(name string, args ...string) string
}

type Context struct {
	Logger  Logger
	AppRoot string
	Port    int
}

func Start(ctx Context) error {
	var pdfController = PdfController{
		Logger: ctx.Logger,
	}

	var routes = Routes()

	var registrationsController = RegistrationsController{
		Pather:      NewPather(ctx.Logger, routes),
		AuthService: (services.AuthService)(nil),
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

type NilLogger struct{}

func (n NilLogger) Log(string, ...interface{}) {}
