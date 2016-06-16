package app

import "net/http"

type SessionsController struct{}

func (s SessionsController) Create(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Set-Cookie", "value")
	rw.WriteHeader(302)
}
