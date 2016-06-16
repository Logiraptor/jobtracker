package app

import "github.com/gorilla/mux"

//go:generate mockgen -source $GOFILE -destination ./mocks/mock_app.go -package mocks

func Routes() *mux.Router {
	routers := mux.NewRouter()
	routers.NewRoute().Name("index").Path("/")
	return routers
}

type Pather interface {
	Path(name string) string
}
