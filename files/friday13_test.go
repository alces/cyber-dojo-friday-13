package friday13

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCountsShouldNotBeEmpty(t *testing.T) {
    assert.NotEmpty(t, count13s(2000, 2000))
}

func TestCountsShouldHaveFridayKey(t *testing.T) {
    assert.Contains(t, count13s(2000, 2000), "Friday")    
}

func TestSumOfCountsShouldBe12(t *testing.T) {
    sum := 0
    
    for _, v := range count13s(2000, 2009) {
        sum += v
    }
    
    assert.Equal(t, 120, sum)
}
        

func TestWeekDay(t *testing.T) {
    assert.Equal(t, "Friday", weekDay(2024, 5, 17))
}
