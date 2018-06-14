package api

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_ApiCalculateDate_Input_StartDay_4_StartMonth_1_StartYear_2018_EndDay_4_EndMonth_1_EndYear_2018_Should_Be_DurationResponse(t *testing.T) {
	url := "/duration?startDay=4&startMonth=1&startYear=2018&endDay=4&endMonth=6&endYear=2018"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expected := `{"days":"152 days"}`

	ApiCalculateDate(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be %s but it got %s", expected, string(actual))
	}
}
