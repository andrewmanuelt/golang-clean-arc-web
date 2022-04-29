package middleware

import (
	"golang-clean-arc-web/helper"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sess := helper.NewSession()

		if sess.GetSession(w, r, "username") == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
