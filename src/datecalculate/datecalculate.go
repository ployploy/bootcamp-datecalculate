package datecalculate

import (
	"fmt"
	"strconv"
	"time"
)

type Duration struct {
	// From    string `json:"from"`
	// To      string `json:"to"`
	Days string `json:"resultDay"`
	// Seconds string `json:"seconds"`
	// Minutes string `json:"minutes"`
	// Hours   string `json:"hours"`
}

func MakeJson(startDate, endDate string) Duration {
	totalDay := CalculateDay(startDate, endDate)
	// hours := TransformDaysToHour(totalDay)
	// minutes := TransformHoursToMinutes(hours)
	// seconds := TransformMinutesToSeconds(minutes)
	// secondsResult := fmt.Sprintf("%s %s", AddComma(int64(seconds)), "seconds")

	return Duration{
		// From:    TransformDateToFullDate(startDay, startMonth, startYear),
		// To:      TransformDateToFullDate(endDay, endMonth, endYear),
		Days: fmt.Sprintf("%d %s", totalDay, "days"),
		// Seconds: secondsResult,
		// Minutes: fmt.Sprintf("%s %s", AddComma(minutes), "minutes"),
		// Hours:   fmt.Sprintf("%s %s", AddComma(hours), "hours"),
	}
}

func CalculateDay(startDate string, endDate string) int {

	startDay, startMonth, startYear := convertDate(startDate)
	endDay, endMonth, endYear := convertDate(endDate)

	startDateAsTime := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endDateAsTime := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	diff := endDateAsTime.Sub(startDateAsTime)

	return (int(diff.Hours()) / 24) + 1
}

func NewDate(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func FormatDate(inputDate string) string {
	day, month, year := convertDate(inputDate)
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return fmt.Sprintf("%v, %v %v %v", date.Weekday(), date.Day(), date.Month(), date.Year())
}

func convertDate(inputDate string) (int, int, int) {
	day, _ := strconv.Atoi(inputDate[0:2])
	month, _ := strconv.Atoi(inputDate[2:4])
	year, _ := strconv.Atoi(inputDate[4:8])

	return day, month, year
}

func RemoveSlashFromStringDate(date string) string {
	var day, month, year int
	fmt.Sscanf(date, "%d/%d/%d", &day, &month, &year)
	return fmt.Sprintf("%02d%02d%d", day, month, year)
}
