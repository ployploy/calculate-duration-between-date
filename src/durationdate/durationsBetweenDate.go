package durationdate

import (
	"strconv"
	"time"
)

func SetFullNameDate(day int, month int, year int) string {

	dateTime := time.Date(
		year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	years := dateTime.Year()
	resultYear := strconv.Itoa(years)

	months := dateTime.Month()
	resultMonth := months.String()

	days := dateTime.Day()
	resultDay := strconv.Itoa(days)

	weekDays := dateTime.Weekday()
	week := weekDays.String()

	fullNameDate := week + " " + resultDay + " " + resultMonth + " " + resultYear
	return fullNameDate

}

func CountBetweenDate(startday int, startmonth int, startyear int, endday int, endmonth int, endyear int) string {
	startDate := time.Date(startyear, time.Month(startmonth), startday, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(endyear, time.Month(endmonth), endday, 0, 0, 0, 0, time.UTC)

	diffDate := endDate.Sub(startDate)
	diffDay := ((diffDate.Hours()) + 24) / 24

	days := strconv.FormatFloat(diffDay, 'f', 0, 64)

	return days

}

func TransferToMinute(days int) string {
	minutes := days * 24 * 60
	return strconv.Itoa(minutes)
}

func TransferToHour(days int) string {
	hours := days * 24
	return strconv.Itoa(hours)

}

func TransferToSecond(days int) string {
	seconds := days * 24 * 60 * 60
	return strconv.Itoa(seconds)
}

func TransferDayToWeek(days int) (string, string) {
	weeks := days / 7
	dayoverweeks := days % 7

	week := strconv.Itoa(weeks)
	dayover := strconv.Itoa(dayoverweeks)

	return week, dayover

}

func TransferToPercentOfYear(days int) string {
	var ResultPercentOfYear float64 = (float64(days) / 365) * 100
	percentofyear := strconv.FormatFloat(ResultPercentOfYear, 'f', 2, 64)
	return percentofyear
}
