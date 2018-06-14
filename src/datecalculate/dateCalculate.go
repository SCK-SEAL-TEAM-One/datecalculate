package datecalculate

import (
	"fmt"
	"time"
)

const HOUR = 24

func NewDate(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func DurationBetweenDate(startDate, endDate time.Time) int {
	diff := endDate.Sub(startDate)
	days := diff.Hours()/HOUR + 1
	return int(days)
}

func FormatDay(days int) string {
	if days < 2 {
		return fmt.Sprintf("%d day", days)
	}
	return fmt.Sprintf("%d days", days)

}
