package apidatecalculate

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_ApiDateCalculate_Input_StartDate_4_1_2018_And_EndDate_4_6_2018_Should_Be_152(t *testing.T) {
	url := "/duration?startdate=4/1/2018&enddate=4/6/2018"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expected := `{"From":"Thursday, 4 January 2018","To":"Monday, 4 June 2018","Days":"152 days","Years":"5 months, 1 day","Seconds":"13,132,800 seconds","Minutes":"218,880 minutes","Hours":"3,648 hours","Weeks":"21 weeks and 5 days","RatioOfYear":"41.64% of 2018"}`

	ApiDateCalculate(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if string(body) != expected {
		t.Errorf("Should be %s but got %s", expected, string(body))
	}
}
