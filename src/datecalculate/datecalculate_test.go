package datecalculate

import (
	"testing"
	"time"
)

func Test_NewDate_Input_4_1_2018_Should_Be_4_1_2018(t *testing.T) {
	day := 4
	month := 1
	year := 2018
	expected := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	actual := NewDate(day, month, year)

	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func Test_NewDate_Input_4_6_2018_Should_Be_4_6_2018(t *testing.T) {
	day := 4
	month := 6
	year := 2018
	expected := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	actual := NewDate(day, month, year)

	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
func Test_MakeJson_Input_StartDate_4_1_2018_EndDate_4_6_2018_Should_Be_Return_DurationResponse_With_152_days(t *testing.T) {
	startDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)
	expected := Duration{
		Days: "152 days",
	}

	actual := MakeJson(startDate, endDate)

	if expected != actual {
		t.Errorf("expected %v but go %v", expected, actual)
	}
}
