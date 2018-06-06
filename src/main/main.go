package main

import (
	"durationdate"
	"net/http"
)

func main() {
	http.HandleFunc("/getDayResult/date/calculatedate", durationdate.Duration)
	http.ListenAndServe(":9000", nil)
}
