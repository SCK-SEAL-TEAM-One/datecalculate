package datecalculate

import (
	"testing"
)

func Test_FormatDate_Input_4_1_2018_Should_Be_Thursday_4_January_2018(t *testing.T) {
	startDay := 4
	startMonth := 1
	startYear := 2018
	expected := "Thursday, 4 January 2018"

	convertDate := FormatDate(startDay, startMonth, startYear)

	if expected != convertDate {
		t.Errorf("Should be %s but got %s", expected, convertDate)
	}

}

func Test_FormatDate_Input_4_6_2018_Should_Be_Thursday_4_june_2018(t *testing.T) {
	endDay := 4
	endMonth := 6
	endYear := 2018
	expected := "Monday, 4 June 2018"

	convertDate := FormatDate(endDay, endMonth, endYear)

	if expected != convertDate {
		t.Errorf("Should be %s but got %s", expected, convertDate)
	}

}

func Test_DayToHour_Input_152_Should_Be_3648(t *testing.T) {
	days := 152
	expected := "3,648 hour"

	convertHour := DayToHour(days)

	if expected != convertHour {
		t.Errorf("Should be %s but got %s", expected, convertHour)
	}

}

func Test_PercentOfYear_Input_152_Should_Be_41_Dot_64(t *testing.T) {
	days := 152
	expected := "41.64% of common year(365 days)"

	daysToPercentOfYear := PercentOfYear(days)

	if expected != daysToPercentOfYear {
		t.Errorf("Should be %s but got %s", expected, daysToPercentOfYear)
	}
}

func Test_DayToMinute_Input_152_Should_Be_218880(t *testing.T) {
	days := 152
	expected := "218,880 minutes"

	convertMinute := DayToMinute(days)

	if expected != convertMinute {
		t.Errorf("Should be %s but got %s", expected, convertMinute)
	}

}

func Test_DayToWeek_Input_152_Should_Be_21_Weeks_And_5_Days(t *testing.T) {
	days := 152
	expected := "21 weeks and 5 days"

	convertDayToWeek := DayToWeek(days)

	if expected != convertDayToWeek {
		t.Errorf("Should be %s but got %s", expected, convertDayToWeek)
	}

}
