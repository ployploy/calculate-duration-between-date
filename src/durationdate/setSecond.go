package durationdate

import "strconv"

// find second
func setSecond(day int) string {
	second := day * 24 * 60 * 60
	return strconv.Itoa(second)
}
