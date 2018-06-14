package api

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_ApiCalculateDate_Input_startDay_4_startMonth_1_startYear_2018_endDay_4_endMonth_6_endYear_2018_Should_Be_Duration_Response(t *testing.T) {
	url := "/duration?startDay=4&startMonth=1&startYear=2018&endDay=4&endMonth=6&endYear=2018"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expected := `{"from":"Thursday, 4 January 2018","to":"Monday, 4 June 2018","days":"152 days","years":"5 months, 1 day","seconds":"13,132,800 seconds","minutes":"218,880 minutes","hours":"3,648 hours","weeks":"21 weeks and 5 days","ratioOfYears":"41.64% of 2018"}`

	ApiCalculateDate(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if string(body) != expected {
		t.Errorf("Should be %s but got %s", expected, string(body))
	}

}
