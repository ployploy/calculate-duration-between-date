package durationdate

import (
	"testing"
)

func Test_Input_152_Transfer_To_Second_Should_Be_13132800(t *testing.T) {
	day := 152
	resultSecond := TransferToSecond(day)
	expected := "13132800"
	if resultSecond != expected {
		t.Errorf("expected %s but %s", expected, resultSecond)
	}

}
