package durationdate

import (
	"testing"
)

func Test_Input_152_Transfer_To_Hour_Should_Be_3648(t *testing.T) {
	day := 152
	resultHours := TransferToHour(day)
	expected := "3648"
	if resultHours != expected {
		t.Errorf("expected %s but %s", expected, resultHours)
	}

}
