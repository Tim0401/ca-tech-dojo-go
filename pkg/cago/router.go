package cago

import (
	"ca-tech-dojo-go/pkg/cago/controller/user"
	"fmt"
	"net/http"
)

// Router ルーティング
func Router(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "hello world.")
		return
	}

	switch r.URL.Path {
	case "/user/create":
		user.Create(w, r)
	case "/user/update":
		user.Update(w, r)
	case "/user/get":
		user.Get(w, r)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Not Found\n")
	}
}
