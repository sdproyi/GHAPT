package components

import (
	"net/http"

	"github.com/a-h/templ"
)

type State struct {
	Open bool
}

func Routes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(Index()).ServeHTTP(w, r)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if r.Header.Get("HX-Request") == "true" {
				templ.Handler(LoginContent()).ServeHTTP(w, r)
			} else {
				templ.Handler(Login()).ServeHTTP(w, r)
			}
		}
	})

	http.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if r.Header.Get("HX-Request") == "true" {
				templ.Handler(SignUpContent()).ServeHTTP(w, r)
			} else {
				templ.Handler(SignUp()).ServeHTTP(w, r)
			}
		}
	})

	http.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ErrorPage()).ServeHTTP(w, r)
	})
}
