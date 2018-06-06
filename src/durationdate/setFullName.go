package durationdate

import (
	"strconv"
	"time"
)

func SetFullName(day int, month int, year int) string {

	then := time.Date(
		year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	years := then.Year()
	var resultyear string
	resultyear = strconv.Itoa(years)

	months := then.Month()
	var resultmonth string
	resultmonth = months.String()

	days := then.Day()
	var resultday string
	resultday = strconv.Itoa(days)

	weekdays := then.Weekday()
	var week string
	week = weekdays.String()

	var results string
	results = week + " " + resultday + " " + resultmonth + " " + resultyear
	return results

}
