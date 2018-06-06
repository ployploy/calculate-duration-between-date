package durationdate

import (
	"testing"
)

func Test_Input_152_Transfer_To_Minute_Should_Be_218880(t *testing.T) {
	day := 152
	resultMinute := TransferToMinute(day)
	expected := "218880"
	if resultMinute != expected {
		t.Errorf("expected %s but %s", expected, resultMinute)
	}

}
