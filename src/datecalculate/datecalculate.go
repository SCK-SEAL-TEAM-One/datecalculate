package datecalculate

import (
	"bytes"
	"fmt"
	"strconv"
	"stringutil"
	"time"
)

const HOUR = 24

func durationBetweenDate(startDay, startMonth, startYear, endDay, endMonth, endYear int) int {
	startDate := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	durationDate := endDate.Sub(startDate)
	days := (durationDate.Hours() / HOUR) + 1

	return int(days)
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
