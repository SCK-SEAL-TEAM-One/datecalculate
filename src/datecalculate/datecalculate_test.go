package datecalculate

import (
	"strconv"
	"testing"
)

func Test_DurationBetweenDate_Input_4_1_2018_And_4_6_2018_Should_Be_152(t *testing.T) {
	startDay := 4
	startMonth := 1
	startYear := 2018
	endDay := 4
	endMonth := 6
	endYear := 2018
	expected := "152"
	actual := DurationBetweenDate(startDay, startMonth, startYear, endDay, endMonth, endYear)
	if expected != strconv.Itoa(actual) {
		t.Errorf("should be %s but it is %d ", expected, actual)
	}
}
