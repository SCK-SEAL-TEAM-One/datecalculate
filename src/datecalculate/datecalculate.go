package datecalculate

import "time"

const HOUR = 24

func durationBetweenDate(startDay, startMonth, startYear, endDay, endMonth, endYear int) int {
	startDate := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	durationDate := endDate.Sub(startDate)
	days := (durationDate.Hours() / HOUR) + 1

	return int(days)
}
