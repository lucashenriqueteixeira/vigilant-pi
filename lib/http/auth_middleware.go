package http

import (
	"net/http"
)

// AuthMiddleware does basic auth when User or Pass is not empty
type AuthMiddleware struct {
	User, Pass string
	h          http.Handler
}

// ServerHTTP ...
func (a *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if a.User != "" || a.Pass != "" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		if user, pass, ok := r.BasicAuth(); !ok || user != a.User || pass != a.Pass {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	a.h.ServeHTTP(w, r)
}

// NewAuthMiddleware ...
func NewAuthMiddleware(handler http.Handler, user, pass string) *AuthMiddleware {
	return &AuthMiddleware{
		User: user,
		Pass: pass,
		h:    handler,
	}
}
