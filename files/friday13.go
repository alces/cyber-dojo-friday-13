package friday13

import (
    "time"
)

func weekDay(year, month, day int) string {
    return time.Time(year, time.Month(month), day, 0, 0, 0, 0).Weekday().String()
}
