package cago

import (
	"fmt"
	"net/http"
)

// Router ルーティング
func Router(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world.")
}
