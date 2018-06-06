package durationdate

import (
	"testing"
)

func Test_Input_Startday_Endday_Count_Between_Date_Should_Be_152(t *testing.T) {
	startday := 4
	startmonth := 1
	startyear := 2018
	endday := 4
	endmonth := 6
	endyear := 2018

	resultDay := CountBetweenDate(startday, startmonth, startyear, endday, endmonth, endyear)

	expected := "152"
	if resultDay != expected {
		t.Errorf("expected %s but %s", expected, resultDay)
	}
}
