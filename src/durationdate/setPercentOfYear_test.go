package durationdate

import (
	"testing"
)

func Test_setPercentOfYear_input_152_Should_be_41_dot_46(t *testing.T) {
	days := 152
	resultPercentOfYear := TransferToPercentOfYear(days)
	expected := "41.64"
	if resultPercentOfYear != expected {
		t.Errorf("expected %s but %s", expected, resultPercentOfYear)
	}

}
