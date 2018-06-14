package datecalculate

import (
	"time"
)

func DurationBetweenDate(startDay, startMonth, startYear, endDay, endMonth, endYear int) int {
	startDate := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	diff := endDate.Sub(startDate)

	durationDate := (diff.Hours() / 24) + 1
	return int(durationDate)
}
