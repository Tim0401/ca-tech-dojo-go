package main

import (
	"ca-tech-dojo-go/internal/cago"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cago.Router(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
