package datecalculate

import (
	"testing"
	"time"
)

func Test_NewtDate_Input_4_1_2018_Should_Be_4_1_2018_Time_date(t *testing.T) {
	startDay := 4
	startMonth := 1
	startYear := 2018
	expected := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)

	DateTime := NewDate(startDay, startMonth, startYear)

	if expected != DateTime {
		t.Errorf("Should be %s but got %s", expected, DateTime)
	}

}

func Test_NewtDate_Input_4_6_2018_Should_Be_4_6_2018_Time_date(t *testing.T) {
	startDay := 4
	startMonth := 6
	startYear := 2018
	expected := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)

	DateTime := NewDate(startDay, startMonth, startYear)

	if expected != DateTime {
		t.Errorf("Should be %s but got %s", expected, DateTime)
	}

}

func Test_MakeJson_Input_startDate_4_1_2018_endDate_4_6_2018_Should_Be_Return_DurationResponse(t *testing.T) {
	startDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)
	expected := DurationResponse{
		From:         "Thursday, 4 January 2018",
		To:           "Monday, 4 June 2018",
		Days:         "152 days",
		Years:        "5 months, 1 day",
		Seconds:      "13,132,800 seconds",
		Minutes:      "218,880 minutes",
		Hours:        "3,648 hours",
		Weeks:        "21 weeks and 5 days",
		RatioOfYears: "41.64% of 2018",
	}
	actualResponse := MakeJson(startDate, endDate)

	if expected != actualResponse {
		t.Errorf("Should be %v but got %v", expected, actualResponse)
	}

}
