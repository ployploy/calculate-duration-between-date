package durationdate

import "testing"

func Test_Input_4_1_2018_SetFullNameDate_Should_Be_Thursday_4_January_2018(t *testing.T) {
	day := 4
	month := 1
	year := 2018
	dateFullName := SetFullNameDate(day, month, year)
	expected := "Thursday 4 January 2018"
	if dateFullName != expected {
		t.Errorf("Should be %s but to be %s", expected, dateFullName)
	}
}
