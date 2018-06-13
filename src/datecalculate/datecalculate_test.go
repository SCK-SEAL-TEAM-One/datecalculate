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
