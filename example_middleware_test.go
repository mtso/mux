package mux_test

import (
	"fmt"
	"net/http"
)

func makeMiddleware(h http.Handler) http.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("%s %s %s", r.Method, r.Code, r.URL)
		h(w, r)
	}
}

func Example_middleware() {

}
