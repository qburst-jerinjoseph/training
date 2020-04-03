package app

import (
	"net/http"
)

func (s *server) middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Info(r.Method, "   ", r.URL, "  ", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
