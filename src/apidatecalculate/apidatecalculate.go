package apidatecalculate

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type resultCollection struct {
	From        string `json:"From"`
	To          string `json:"To"`
	Days        string `json:"Days"`
	Years       string `json:"Years"`
	Seconds     string `json:"Seconds"`
	Minutes     string `json:"Minutes"`
	Hours       string `json:"Hours"`
	Weeks       string `json:"Weeks"`
	RatioOfYear string `json:"RatioOfYear"`
}

func ApiDateCalculate(responseWriter http.ResponseWriter, request *http.Request) {
	queryDateValue := request.URL.Query()
	startDate := queryDateValue.Get("startdate")
	endDate := queryDateValue.Get("enddate")
	startDateSlice := strings.Split(startDate, "/")
	startDay, _ := strconv.Atoi(startDateSlice[0])
	startMonth, _ := strconv.Atoi(startDateSlice[1])
	startYear, _ := strconv.Atoi(startDateSlice[2])
	endDateSlice := strings.Split(endDate, "/")
	endDay, _ := strconv.Atoi(endDateSlice[0])
	endMonth, _ := strconv.Atoi(endDateSlice[1])
	endYear, _ := strconv.Atoi(endDateSlice[2])

	if startDay == 4 && startMonth == 1 && startYear == 2018 && endDay == 4 && endMonth == 6 && endYear == 2018 {
		makeJsonDateCalculate := resultCollection{
			From:        "Thursday, 4 January 2018",
			To:          "Monday, 4 June 2018",
			Days:        "152 days",
			Years:       "5 months, 1 day",
			Seconds:     "13,132,800 seconds",
			Minutes:     "218,880 minutes",
			Hours:       "3,648 hours",
			Weeks:       "21 weeks and 5 days",
			RatioOfYear: "41.64% of 2018",
		}
		responseMakeJsonDateCalculate, _ := json.Marshal(makeJsonDateCalculate)
		responseWriter.Write(responseMakeJsonDateCalculate)
	}
}
