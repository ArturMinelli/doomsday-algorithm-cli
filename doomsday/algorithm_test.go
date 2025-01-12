package doomsday

import (
	"math/rand"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		min := time.Date(1700, 1, 1, 0, 0, 0, 0, time.UTC)
		max := time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC)

		date := time.Date(rand.Intn(max.Year()-min.Year()+1)+min.Year(), time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.UTC)

		result := Run(date)

		if int(date.Weekday()) != result.Weekday {
			t.Errorf(
				"Run() for %v = weekday %d; want weekday %d",
				date.Format("2006-01-02"),
				result.Weekday,
				int(date.Weekday()),
			)
		}
	}
}
