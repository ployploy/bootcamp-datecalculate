package api

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_ApiCalculateDate_Input_StartDay_7_StartMonth_3_StartYear_1977_EndDay_12_EndMonth_7_EndYear_1977_Should_Be_DurationResponse(t *testing.T) {
	url := "/duration?startDay=7&startMonth=3&startYear=1977&endDay=12&endMonth=7&endYear=2018"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expected := `{"days":"15103 days"}`

	ApiCalculateDate(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be %s but it got %s", expected, string(actual))
	}
}
