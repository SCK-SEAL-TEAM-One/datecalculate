package datecalculate

import "time"

const HOUR = 24

func NewDate(day, month, year int) time.Time {

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func DurationBetweenDate(startDay, startMonth, startYear, endDay, endMonth, endYear int) int {
	startDate := NewDate(startDay, startMonth, startYear)
	endDate := NewDate(endDay, endMonth, endYear)
	diff := endDate.Sub(startDate)
	days := diff.Hours()/HOUR + 1

	return int(days)
}
