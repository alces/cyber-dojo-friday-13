package friday13

import (
    "time"
)

func count13s(firstYear, lastYear int) map[string]int {
    result := make(map[string]int)
    
    for y := firstYear; y <= lastYear; y++ {
        for m := 1; m < 13; m++ {
            result[weekDay(year, m, 13)]++
        }
    }
    
    return result
}

func weekDay(year, month, day int) string {
    date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    return date.Weekday().String()
}
