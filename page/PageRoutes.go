package page

import (
	"net/http"

	"github.com/a-h/templ"
)

type State struct {
	Open bool
}

func PageRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(IndexPage()).ServeHTTP(w, r)
	})

}
