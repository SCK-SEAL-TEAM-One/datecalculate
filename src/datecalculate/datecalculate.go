package datecalculate

import "fmt"

func FormatDay(days int) string {
	if days < 2 {
		return fmt.Sprintf("%d day", days)
	}
	return fmt.Sprintf("%d days", days)

}
