package web

import (
	"html/template"
	"net/http"
)

type View interface {
	Render(rw http.ResponseWriter, name string, data interface{}) error
}

type TemplateView struct {
	templates *template.Template
}

func NewTemplateView(tmpl *template.Template) TemplateView {
	return TemplateView{
		templates: tmpl,
	}
}

func (t TemplateView) Render(rw http.ResponseWriter, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(rw, name, data)
}
