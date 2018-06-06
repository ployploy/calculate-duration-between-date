package durationdate

import (
	"testing"
)

func Test_setPercentOfYear_input_152_Should_be_41_dot_46(t *testing.T) {
	day := 152
	resultPercentOfYear := TransferToPercentOfYear(day)
	expected := "41.64"
	if resultPercentOfYear != expected {
		t.Errorf("expected %s but %s", expected, resultPercentOfYear)
	}

}
