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

	DateTime := NewDate(day, month, year)

	if expected != DateTime {
		t.Errorf("Should be %s but got %s", expected, DateTime)
	}

}

func Test_NewtDate_Input_4_6_2018_Should_Be_4_6_2018_Time_date(t *testing.T) {
	day := 4
	month := 6
	year := 2018
	expected := time.Date(2018, 6, 4, 0, 0, 0, 0, time.UTC)

	DateTime := NewDate(day, month, year)

	if expected != DateTime {
		t.Errorf("Should be %s but got %s", expected, DateTime)
	}

}

func Test_DurationBetweenDate_Input_4_1_2018_And_4_6_2018_Should_Be_152(t *testing.T) {
	startDay := 4
	startMonth := 1
	startYear := 2018
	endDay := 4
	endMonth := 6
	endYear := 2018
	expected := 152

	actual := DurationBetweenDate(startDay, startMonth, startYear, endDay, endMonth, endYear)

	if expected != actual {
		t.Errorf("Should be %d but got %d", expected, actual)
	}

}

func Test_DurationBetweenDate_Input_27_12_1994_And_4_6_2018_Should_Be_152(t *testing.T) {
	startDay := 27
	startMonth := 12
	startYear := 1994
	endDay := 4
	endMonth := 6
	endYear := 2018
	expected := 8561

	actual := DurationBetweenDate(startDay, startMonth, startYear, endDay, endMonth, endYear)

	if expected != actual {
		t.Errorf("Should be %d but got %d", expected, actual)
	}

}
