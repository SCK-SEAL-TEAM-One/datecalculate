package datecalculate

import (
	"testing"
	"time"
)

func Test_DurationBetweenDate_Input_4_1_2018_And_4_6_2018_Should_Be_152(t *testing.T) {
	startDate := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)
	expected := 152

	actual := DurationBetweenDate(startDate, endDate)

	if expected != actual {
		t.Errorf("should be %d but it is %d ", expected, actual)

	}
}
func Test_FormatDay_Input_152_Should_Be_152_Days(t *testing.T) {
	days := 152
	expected := "152 days"

	actual := FormatDay(days)

	if expected != actual {
		t.Errorf("Should be %s but got %s", expected, actual)
	}

}

func Test_FormatDay_Input_1_Should_Be_1_Day(t *testing.T) {
	days := 1
	expected := "1 day"

	actual := FormatDay(days)

	if expected != actual {
		t.Errorf("Should be %s but got %s", expected, actual)
	}

}

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
