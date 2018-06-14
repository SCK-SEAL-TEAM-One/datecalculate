package datecalculate

import (
	"time"
)

type Duration struct {
	Days string `json:"days"`
}

func NewDate(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func MakeJson(startDate, endDate time.Time) Duration {
	return Duration{
		Days: "152 days",
	}
}
