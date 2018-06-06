package durationdate

import (
	"encoding/json"
	"log"
	"net/http"
)

type dayResult struct {
	DayNameFrom   string `json:"daynamefrom"`
	DayFrom       string `json:"dayfrom"`
	MonthNameFrom string `json:"monthnamefrom"`
	YearFrom      string `json:"yearfrom"`
	DayNameTo     string `json:"daynameto"`
	DayTo         string `json:"dayto"`
	MonthNameTo   string `json:"monthnameto"`
	YearTo        string `json:"yearto"`
	ResultDay     string `json:"resultday"`
	ResultMonth   string `json:"resultmonth"`
	ResultYear    string `json:"resultyear"`
	ResultOverDay string `json:"resultoverday"`
	Second        string `json:"second"`
	Minute        string `json:"minute"`
	Hours         string `json:"hours"`
	Week          string `json:"week"`
	DayOverOfWeek string `json:"dayoverofweek"`
	Percent       string `json:"percent"`
}
type statusCode struct {
	Status string `json:"status"`
}

func GetDayResult(w http.ResponseWriter, r *http.Request) {
	dayResultStruct := dayResult{
		DayNameFrom:   "Thursday",
		DayFrom:       "4",
		MonthNameFrom: "January",
		YearFrom:      "2018",
		DayNameTo:     "Monday",
		DayTo:         "4",
		MonthNameTo:   "June",
		YearTo:        "2018",
		ResultDay:     "152",
		ResultYear:    "0",
		ResultMonth:   "5",
		ResultOverDay: "1",
		Second:        "13,132,800",
		Minute:        "218,880",
		Hours:         "3,648",
		Week:          "21",
		DayOverOfWeek: "5",
		Percent:       "41.64",
	}
	jsonMarshallDayResultStruct, err := json.Marshal(dayResultStruct)
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonMarshallDayResultStruct)

}
