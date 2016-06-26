package web

import "github.com/gorilla/mux"

type Pather interface {
	Path(name string, args ...string) string
}

type routePather struct {
	Logger
	routes *mux.Router
}

func (r routePather) Path(name string, args ...string) string {
	if route := r.routes.Get(name); route != nil {
		url, err := route.URL(args...)
		if err == nil {
			return url.String()
		}
	}
	r.Log("undefined route: %s %s", name, args)
	return "/"
}

func NewPather(logger Logger, routes *mux.Router) routePather {
	return routePather{
		Logger: logger,
		routes: routes,
	}
}
