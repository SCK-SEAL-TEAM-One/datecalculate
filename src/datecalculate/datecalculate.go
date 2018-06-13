package datecalculate

import (
	"fmt"
	"strconv"
	"time"
)

func FormatDate(day, month, year int) string {
	fullNameDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	convertWeekDay := fullNameDate.Weekday()
	convertDay := fullNameDate.Day()
	convertMonth := fullNameDate.Month()
	convertYear := fullNameDate.Year()

	return fmt.Sprintf("%s, %s %s %s", convertWeekDay.String(), strconv.Itoa(convertDay), convertMonth.String(), strconv.Itoa(convertYear))
}

func DayToHour(days int) string {
	hour := days * 24

	return fmt.Sprintf("%s hours", strconv.Itoa(hour))
}

func PercentOfYear(days int) string {
	numberPercentOfYear := (float64(days) / 365) * 100

	return fmt.Sprintf("%s of common year(365 days)", strconv.FormatFloat(numberPercentOfYear, 'f', 2, 64))
}

func DayToMinute(days int) string {
	minute := days * 24 * 60

	return fmt.Sprintf("%s minutes", strconv.Itoa(minute))
}

func DayToWeek(days int) string {
	week := days / 7
	overday := days % 7

	return fmt.Sprintf("%s weeks and %s days", strconv.Itoa(week), strconv.Itoa(overday))
}
