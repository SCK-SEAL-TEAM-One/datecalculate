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
