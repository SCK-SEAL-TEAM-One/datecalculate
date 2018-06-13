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
