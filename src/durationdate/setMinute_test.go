package durationdate

import (
	"testing"
)

func Test_setMinute_input_152_Should_be_218880(t *testing.T) {
	day := 152
	resultMinute := TransferToMinute(day)
	expected := "218880"
	if resultMinute != expected {
		t.Errorf("expected %s but %s", expected, resultMinute)
	}

}
