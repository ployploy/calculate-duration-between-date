package durationdate

import "strconv"

func setSecond(day int) string {
	second := day * 24 * 60 * 60
	return strconv.Itoa(second)
}
