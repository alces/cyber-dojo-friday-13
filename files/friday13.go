package friday13

import (
    "time"
)

func count13s(year int) map[string]int {
    return make(map[string]int)   
}

func weekDay(year, month, day int) string {
    date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    return date.Weekday().String()
}
