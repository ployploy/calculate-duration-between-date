package durationdate

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetStartDateAndEndDate_input_start_date_4_1_2018_and_end_date_4_6_2018_should_be_4_1_2018_and_4_6_2018(t *testing.T) {
	url := "http://localhost:9000/getDayResult/date/calculatedate?startdate=4/1/2018&enddate=4/6/2018"
	expected := dateInput{DayFrom: "4", MonthFrom: "1", YearFrom: "2018", DayTo: "4", MonthTo: "6", YearTo: "2018"}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("error")
	}

	startDate := GetStartDateAndEndDate(req)

	if startDate != expected {
		t.Errorf("should be %s but got %s", expected, startDate)
	}

}

func Test_Duration_input_start_date_4_1_2018_and_end_date_4_6_2018_should_be_4_1_2018_and_4_6_2018(t *testing.T) {
	url := "http://localhost:9000/getDayResult/date/calculatedate?startdate=4/1/2018&enddate=4/6/2018"
	req := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	expected := `{"daynamefrom":"Thursday","dayfrom":"4","monthnamefrom":"January","yearfrom":"2018","daynameto":"Monday","dayto":"4","monthnameto":"June","yearto":"2018","resultday":"152","resultmonth":"5","resultyear":"0","resultoverday":"1","second":"13,132,800","minute":"218,880","hours":"3,648","week":"21","dayoverofweek":"5","percent":"41.64"}`

	Duration(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != expected {
		t.Errorf("should be %s but got %s", expected, string(body))
	}

}
