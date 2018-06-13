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

func FormatDay(days int) string {
	if days < 2 {
		return fmt.Sprintf("%d day", days)
	}
	return fmt.Sprintf("%d days", days)
}