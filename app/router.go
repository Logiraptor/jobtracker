package app

import "github.com/gorilla/mux"

func Routes() *mux.Router {
	routers := mux.NewRouter()
	routers.NewRoute().Path("/login").Name("login")
	routers.NewRoute().Path("/pdf").Name("generate_pdf")
	routers.NewRoute().Path("/sign_up").Methods("POST").Name("sign_up")

	routers.NewRoute().PathPrefix("/").Name("index")
	return routers
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

func NewPather(logger Logger, routes *mux.Router) Pather {
	return routePather{
		Logger: logger,
		routes: routes,
	}
}
