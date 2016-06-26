package web

import "github.com/gorilla/mux"

func Routes() *mux.Router {
	routers := mux.NewRouter()
	routers.NewRoute().Path("/login").Name("login")
	routers.NewRoute().Path("/pdf").Name("generate_pdf")
	routers.NewRoute().Path("/sign_up").Methods("POST").Name("sign_up")

	routers.NewRoute().PathPrefix("/").Name("index")
	return routers
}
