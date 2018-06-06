package durationdate

import (
	"testing"
)

func Test_setWeeks_input_152_Should_be_21Weeks(t *testing.T) {
	day := 152
	resultWeeks, dayover := setWeeks(day)
	expected1 := "21"
	expected2 := "5"
	if resultWeeks != expected1 {
		t.Errorf("expected %s but %s", expected1, resultWeeks)

	}
	if dayover != expected2 {
		t.Errorf("expected %s but %s", expected2, dayover)
	}
}
