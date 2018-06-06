package durationdate

import (
	"testing"
)

func Test_setSecond_input_152_Should_be_13132800(t *testing.T) {
	day := 152
	resultSecond := setSecond(day)
	expected := "13132800"
	if resultSecond != expected {
		t.Errorf("expected %s but %s", expected, resultSecond)
	}

}
