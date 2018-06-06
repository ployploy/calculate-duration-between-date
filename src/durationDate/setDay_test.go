package durationDate

import (
	"testing"
)

func Test_setDay_input_startday_endday_Should_be_152(t *testing.T) {
	startday := 4
	startmonth := 1
	startyear := 2018
	endday := 4
	endmonth := 6
	endyear := 2018

	resultDay := SetDay(startday, startmonth, startyear, endday, endmonth, endyear)

	expected := "152"
	if resultDay != expected {
		t.Errorf("expected %s but %s", expected, resultDay)
	}
}
