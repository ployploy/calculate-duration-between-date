package durationdate

import (
	"testing"
)

func Test_Input_152_Transfer_Day_To_Week_Should_Be_21_Weeks_5_Days(t *testing.T) {
	day := 152
	resultWeeks, dayover := TransferDayToWeek(day)
	expected1 := "21"
	expected2 := "5"
	if resultWeeks != expected1 {
		t.Errorf("expected %s but %s", expected1, resultWeeks)

	}
	if dayover != expected2 {
		t.Errorf("expected %s but %s", expected2, dayover)
	}
}
