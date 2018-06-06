package durationdate

import "strconv"

func setWeeks(day int) (string, string) {
	weeks := day / 7
	dayoverweeks := day % 7

	week := strconv.Itoa(weeks)
	dayover := strconv.Itoa(dayoverweeks)

	return week, dayover

}
