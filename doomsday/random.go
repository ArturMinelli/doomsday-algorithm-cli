package doomsday

import (
	"math/rand"
	"time"
)

func NewRandomDate() time.Time {
	min := time.Date(1700, 1, 1, 0, 0, 0, 0, time.UTC)
	max := time.Date(2099, 12, 31, 0, 0, 0, 0, time.UTC)

	return time.Date(
		rand.Intn(max.Year()-min.Year()+1)+min.Year(),
		time.Month(rand.Intn(12)+1),
		rand.Intn(28)+1,
		0,
		0,
		0,
		0,
		time.UTC,
	)
}
