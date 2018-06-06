package main

import (
	"durationdate"
	"net/http"
)

func main() {
	http.HandleFunc("/getDayResult/date/calculatedate", durationdate.GetDayResult)
	durationdate.DurationBetweenDate()
	http.ListenAndServe(":9000", nil)
}
