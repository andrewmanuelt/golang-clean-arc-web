package middleware

import (
	"net/http"
)

func StaticfileMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/asset" {
			http.Redirect(w, r, r.Header.Get("referer"), http.StatusFound)
		}

		next.ServeHTTP(w, r)
	})
}
