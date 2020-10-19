package main

import (
	"ca-tech-dojo-go/pkg/cago"
	"net/http"
)

func main() {
	http.HandleFunc("/", cago.Router)
	http.ListenAndServe(":8080", nil)
}
