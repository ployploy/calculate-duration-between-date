package durationdate

import "strconv"

func setMinute(day int) string {
	minute := day * 24 * 60
	return strconv.Itoa(minute)
}
