package friday13

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestWeekDay(t *testing.T) {
    assert.Equal(t, "Friday", weekDay(2024, 5, 17))
}
