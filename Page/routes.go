package page

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


	http.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(ErrorPage()).ServeHTTP(w, r)
	})
}
