package datecalculate

import (
	"testing"
	"time"
)

func Test_MakeJson_Input_StartDate_7_3_1977_EndDate_12_7_2018_Should_Be_Return_DurationResponse_With_15103_days(t *testing.T) {
	startdate := "07031977"
	enddate := "12072018"
	expected := Duration{
		Days: "15103 days",
	}

	actual := MakeJson(startdate, enddate)

	if expected != actual {
		t.Errorf("expected %v but go %v", expected, actual)
	}
}

func Test_CalculateDay_Input_7_3_1977_And_12_7_2018_Should_Be_15103(t *testing.T) {
	startdate := "07031977"
	enddate := "12072018"
	expected := 15103

	actual := CalculateDay(startdate, enddate)

	if expected != actual {
		t.Errorf("Expected %d Actual %d", expected, actual)
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

func Test_RemoveSlashFromStringDate_Input_4Slash1Slash2018_ShouldBe_04012018(t *testing.T) {
	date := "7/3/1977"
	expected := "07031977"
	actual := RemoveSlashFromStringDate(date)
	if actual != expected {
		t.Errorf("Expected: %s but got %s", expected, actual)
	}
}

func Test_FormatDate_Input_07031977_ShouldBe_Monday_7_March_1977(t *testing.T) {
	expected := "Monday, 7 March 1977"
	startDate := "07031977"

	actualDate := FormatDate(startDate)

	if actualDate != expected {
		t.Errorf("expected %s but got %s", expected, actualDate)
	}
}
