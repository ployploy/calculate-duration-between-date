package durationdate

import (
	"testing"
)

func Test_setHours_input_152_Should_be_3648(t *testing.T) {
	day := 152
	resultHours := TransferToHour(day)
	expected := "3648"
	if resultHours != expected {
		t.Errorf("expected %s but %s", expected, resultHours)
	}

}
