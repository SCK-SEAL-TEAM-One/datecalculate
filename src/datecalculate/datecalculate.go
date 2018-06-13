package datecalculate

import (
	"bytes"
	"fmt"
	"strconv"
	"stringutil"
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

func DayToSecond(days int) string {
	seconds := days * 24 * 60 * 60
	return fmt.Sprintf("%s seconds", AddComma(seconds))
}

func AddComma(number int) string {
	num := strconv.Itoa(number)
	var buffer bytes.Buffer
	reverseNum := stringutil.Reverse(num)
	count := 0
	for i := 0; i < len(num); i++ {
		count++
		buffer.WriteString(reverseNum[i : i+1])

		if count == 3 && i != len(num)-1 {
			buffer.WriteString(",")
			count = 0
		}
	}

	return stringutil.Reverse(buffer.String())
}
func FormatDay(days int) string {
	if days < 2 {
		return fmt.Sprintf("%d day", days)
	}
	return fmt.Sprintf("%d days", days)
}

func DurationBetweenDate(startDay, startMonth, startYear, endDay, endMonth, endYear int) int {
	startDate := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	diff := endDate.Sub(startDate)

	durationDate := (diff.Hours() / 24) + 1

	return int(durationDate)

}
