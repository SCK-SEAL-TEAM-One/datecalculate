package datecalculate

import (
	"testing"
)

func Test_durationBetweenDate_input_4_1_2018_and_4_6_2018_should_be_152(t *testing.T) {
	startDay := 4
	startMonth := 1
	startYear := 2018
	endDay := 4
	endMonth := 6
	endYear := 2018
	expected := 152

	days := durationBetweenDate(startDay, startMonth, startYear, endDay, endMonth, endYear)

	if expected != days {
		t.Errorf("Should be %d but got %d", expected, days)
	}
}
