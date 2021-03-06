package middlewares

import (
	"log"
	"maieutic-server/env"
	"maieutic-server/utils"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func logger (h http.Handler) http.Handler {
	PADDING := "================================"
	consumeBody := env.GetOr("GO_ENV", "development") == "development"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rc, err := httputil.DumpRequest(r, consumeBody)
		utils.CheckErr(err)
		log.Printf("\n%s\n%s requested %s:\n    %s\n%s", PADDING, r.RemoteAddr, r.URL, string(rc), PADDING)

		rw := httptest.NewRecorder()
		h.ServeHTTP(rw, r)

		wc, err := httputil.DumpResponse(rw.Result(), consumeBody)
		utils.CheckErr(err)

		log.Printf("\n%s\nfrank replied to %s:\n    %s\n%s", PADDING, r.RemoteAddr, string(wc), PADDING)

		for k, v := range rw.Header() {
			w.Header()[k] = v
		}

		body := rw.Body.Bytes()
		for key := range rw.HeaderMap {
			w.Header().Set(key, w.Header().Get(key))
		}
		w.WriteHeader(rw.Result().StatusCode)
		w.Write(body)
	})
}