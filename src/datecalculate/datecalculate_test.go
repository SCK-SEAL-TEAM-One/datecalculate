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

func Test_DayToSecond_Input_152_Should_Be_13_comma_132_comma_800_seconds(t *testing.T) {
	days := 152
	expected := "13,132,800 seconds"

	actual := DayToSecond(days)

	if actual != expected {
		t.Errorf("expected %s but it is %s", expected, actual)
	}
}

func Test_AddComma_Input_1000_should_be_1_comma_000(t *testing.T) {
	number := 1000
	expected := "1,000"

	actual := AddComma(number)

	if actual != expected {
		t.Errorf("expected %s but it is %s", expected, actual)
	}
}
func Test_AddComma_Input_1000000_should_be_1_comma_000_comma_000(t *testing.T) {
	number := 1000000
	expected := "1,000,000"

	actual := AddComma(number)

	if actual != expected {
		t.Errorf("expected %s but it is %s", expected, actual)
	}
}
func Test_AddComma_Input_1000000000_should_be_1_comma_000_comma_000_comma_000(t *testing.T) {
	number := 1000000000
	expected := "1,000,000,000"

	actual := AddComma(number)

	if actual != expected {
		t.Errorf("expected %s but it is %s", expected, actual)
	}
}
func Test_FormatDay_Input_152_Should_Be_152_Days(t *testing.T) {
	days := 152
	expected := "152 days"

	formatedDay := FormatDay(days)

	if expected != formatedDay {
		t.Errorf("Should be %s but got %s", expected, formatedDay)
	}
}
