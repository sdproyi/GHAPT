package page

import (
	"net/http"

	"github.com/a-h/templ"
)

func PageRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(IndexPage()).ServeHTTP(w, r)
	})

}
