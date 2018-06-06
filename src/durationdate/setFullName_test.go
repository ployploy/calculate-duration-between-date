package durationdate

import "testing"

func Test_SetFullName_input_4_1_2018_should_be_Thursday_4_January_2018(t *testing.T) {
	day := 4
	month := 1
	year := 2018
	dateFullName := SetFullName(day, month, year)
	expected := "Thursday 4 January 2018"
	if dateFullName != expected {
		t.Errorf("Should be %s but to be %s", expected, dateFullName)
	}
}
