package datecalculate

import (
	"testing"
	"time"
)

func Test_NewtDate_Input_4_1_2018_Should_Be_4_1_2018_Time_date(t *testing.T) {
	day := 4
	month := 1
	year := 2018
	expected := time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC)

	actual := NewDate(day, month, year)

	if expected != actual {
		t.Errorf("Should be %s but got %s", expected, actual)
	}

}

func Test_NewtDate_Input_4_6_2018_Should_Be_4_6_2018_Time_date(t *testing.T) {
	day := 4
	month := 6
	year := 2018
	expected := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)

	actual := NewDate(day, month, year)

	if expected != actual {
		t.Errorf("Should be %s but got %s", expected, actual)
	}

}

func Test_DurationBetweenDate_Input_4_1_2018_And_4_6_2018_Should_Be_152(t *testing.T) {
	startDate := NewDate(4, 1, 2018)
	endDate := NewDate(4, 6, 2018)
	expected := 152

	actual := DurationBetweenDate(startDate, endDate)

	if expected != actual {
		t.Errorf("Should be %d but got %d", expected, actual)
	}

}

func Test_DurationBetweenDate_Input_27_12_1994_And_4_6_2018_Should_Be_152(t *testing.T) {
	startDate := NewDate(27, 12, 1994)
	endDate := NewDate(4, 6, 2018)
	expected := 8561

	actual := DurationBetweenDate(startDate, endDate)

	if expected != actual {
		t.Errorf("Should be %d but got %d", expected, actual)
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
