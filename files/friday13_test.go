package friday13

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCountsShouldNotBeEmpty(t *testing.T) {
    assert.NotEmpty(t, count13s(2000))
}

func TestWeekDay(t *testing.T) {
    assert.Equal(t, "Friday", weekDay(2024, 5, 17))
}
