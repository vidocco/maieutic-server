package middlewares

import "net/http"

func auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		// TODO: add auth middleware here
		h.ServeHTTP(w, r)
	})
}
