package user

import (
	"fmt"
	"net/http"
)

// Create POST Only
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "/user/create")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprint(w, "Method not allowed.\n")
}

// Get GET Only
func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "/user/get")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprint(w, "Method not allowed.\n")
}

// Update PUT
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "/user/update")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprint(w, "Method not allowed.\n")
}
