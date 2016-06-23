package app

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type TemplateController struct {
	Template   *template.Template
	AppContext Context
}

func (t TemplateController) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	switch {
	case path == "/":
		t.AppContext.Logger.Log("rendering template: %s", t.Template.Name())
		t.Template.Execute(rw, t)
	default:
		base := filepath.Base(path)
		if filepath.Ext(base) == "" {
			base += ".html"
		}
		if tmpl := t.Template.Lookup(base); tmpl != nil {
			t.AppContext.Logger.Log("rendering template: %s", base)
			tmpl.Execute(rw, t)
		} else {
			t.AppContext.Logger.Log("could not find template: %s", base)
			http.NotFound(rw, req)
		}
	}
}
