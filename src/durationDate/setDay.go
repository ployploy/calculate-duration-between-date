package durationDate

import (
	"strconv"
	"time"
)

func SetDay(startday int, startmonth int, startyear int, endday int, endmonth int, endyear int) string {
	startDate := time.Date(startyear, time.Month(startmonth), startday, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(endyear, time.Month(endmonth), endday, 0, 0, 0, 0, time.UTC)

	diff := endDate.Sub(startDate)
	diffDay := ((diff.Hours()) + 24) / 24

	Day := strconv.FormatFloat(diffDay, 'f', 0, 64)

	return Day

}
