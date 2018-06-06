package durationdate

import "strconv"

func setHours(day int) string {
	hours := day * 24
	return strconv.Itoa(hours)

}
