package authentication

import "net/http"

func RequireAuthentication(sessions HTTPSessionTracker, authedHandler, unauthedHandler http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if _, ok := sessions.CurrentUser(req); ok {
			authedHandler.ServeHTTP(rw, req)
			return
		}
		unauthedHandler.ServeHTTP(rw, req)
	}
}
