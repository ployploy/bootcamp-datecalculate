package datecalculate

import (
	"fmt"
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

func MakeJson(startDate, endDate time.Time) Duration {
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

func CalculateDay(startDate, endDate time.Time) int {
	diff := endDate.Sub(startDate)

	durationDate := (diff.Hours() / 24) + 1
	return int(durationDate)
}

func NewDate(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
