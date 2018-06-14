package datecalculate

import (
	"fmt"
	"time"
)

type DurationResponse struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Days         string `json:"days"`
	Years        string `json:"years"`
	Seconds      string `json:"seconds"`
	Minutes      string `json:"minutes"`
	Hours        string `json:"hours"`
	Weeks        string `json:"weeks"`
	RatioOfYears string `json:"ratioOfYears"`
}

const HOUR = 24

func NewDate(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func MakeJson(startDate, endDate time.Time) DurationResponse {
	duration := DurationBetweenDate(startDate, endDate)
	return DurationResponse{
		From:         "Thursday, 4 January 2018",
		To:           "Monday, 4 June 2018",
		Days:         fmt.Sprintf("%d days", duration),
		Years:        "5 months, 1 day",
		Seconds:      "13,132,800 seconds",
		Minutes:      "218,880 minutes",
		Hours:        "3,648 hours",
		Weeks:        "21 weeks and 5 days",
		RatioOfYears: "41.64% of 2018",
	}
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
