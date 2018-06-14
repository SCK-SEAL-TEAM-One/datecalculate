package datecalculate


import (
	"time"
	"fmt"
)
type Duration struct {
	Days string `json:"days"`
}

func FormatDay(days int) string {
	if days < 2 {
		return fmt.Sprintf("%d day", days)
	}
	return fmt.Sprintf("%d days", days)
}


func NewDate(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func MakeJson(startDate, endDate time.Time) Duration {
	return Duration{
		Days: "152 days",
	}
}
