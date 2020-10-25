package main

import (
	"ca-tech-dojo-go/pkg/cago"
	"net/http"
	"time"
)

const location = "Asia/Tokyo"

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	http.HandleFunc("/", cago.Router)
	http.ListenAndServe(":8080", nil)
}
