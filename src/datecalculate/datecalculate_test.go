package datecalculate

import (
	"testing"
	"time"
)

func Test_MakeJson_Input_StartDate_7_3_1977_EndDate_12_7_2018_Should_Be_Return_DurationResponse_With_152_days(t *testing.T) {
	startDate := time.Date(1977, 3, 7, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 7, 12, 0, 0, 0, 0, time.UTC)
	expected := Duration{
		Days: "15103 days",
	}

	actual := MakeJson(startDate, endDate)

	if expected != actual {
		t.Errorf("expected %v but go %v", expected, actual)
	}
}

func Test_CalculateDay_Input_7_3_1977_And_12_7_2018_Should_Be_152(t *testing.T) {
	startDate := time.Date(1977, 3, 7, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 7, 12, 0, 0, 0, 0, time.UTC)
	expected := 15103

	actual := CalculateDay(startDate, endDate)

	if expected != actual {
		t.Errorf("should be %d but it is %d ", expected, actual)

	}
}

func Test_NewDate_Input_7_3_1977_Should_Be_7_3_2018(t *testing.T) {
	day := 7
	month := 3
	year := 1977
	expected := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	actual := NewDate(day, month, year)

	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
