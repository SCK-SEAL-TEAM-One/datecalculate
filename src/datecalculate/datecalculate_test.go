package datecalculate

import "testing"

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