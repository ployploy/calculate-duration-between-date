package durationdate

import "strconv"

func TransferToPercentOfYear(day int) string {
	var ResultPercentOfYear float64 = (float64(day) / 365) * 100
	percentofyear := strconv.FormatFloat(ResultPercentOfYear, 'f', 2, 64)
	return percentofyear
}
