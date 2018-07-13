package api

import (
	"datecalculate"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_ApiCalculateDate_Input_StartDay_7_StartMonth_3_StartYear_1977_EndDay_12_EndMonth_7_EndYear_1977_Should_Be_DurationResponse(t *testing.T) {
	url := "localhost:3000/date/duration?startDate=7/3/1977&endDate=12/7/2018"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expected := `{"resultDay":"15103 days"}`

	ApiCalculateDate(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be %s but it got %s", expected, string(actual))
	}
}

func Test(t *testing.T) {
	startDay := "7/3/1977"
	accualRemove := datecalculate.RemoveSlashFromStringDate(startDay)
	expected := "07031977"

	endDays := "12/7/2018"
	accualEndDateWithoutSlash := datecalculate.RemoveSlashFromStringDate(endDays)
	expected2 := "12072018"

	startDate := datecalculate.FormatDate(accualRemove)
	expected3 := "Monday, 7 March 1977"

	endDate := datecalculate.FormatDate(accualEndDateWithoutSlash)
	expected4 := "Thursday, 12 July 2018"

	durationResponse, _ := json.Marshal(datecalculate.MakeJson(expected, expected2))
	expected5 := `{"days":"15103 days"}`
	if accualRemove != expected {
		t.Errorf("%s but got %s", accualRemove, expected)
	}

	if accualEndDateWithoutSlash != expected2 {
		t.Errorf("%s but got %s", accualEndDateWithoutSlash, expected2)
	}

	if startDate != expected3 {
		t.Errorf("%s but got %s", startDate, expected3)
	}
	if endDate != expected4 {
		t.Errorf("%s but got %s", endDate, expected4)
	}
	if string(durationResponse) != expected5 {
		t.Errorf("%s but got %s", durationResponse, expected5)
	}
}
