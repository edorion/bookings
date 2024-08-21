package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf is a middleware that implements CSRF protection
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.TLSenabled,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// LoadAndSave is a middleware that loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
