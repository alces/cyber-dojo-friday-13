package friday13

import (
    "time"
)

func weekDay(year, month, day int) string {
    date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    return date.Weekday().String()
}
