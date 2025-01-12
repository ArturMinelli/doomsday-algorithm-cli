package doomsday

var (
	centuryCodes = map[int]int{
		1700: 0,
		1800: 5,
		1900: 3,
		2000: 2,
	}
	monthDoomsdays = map[int]int{
		1:  3,
		2:  28,
		3:  14,
		4:  4,
		5:  9,
		6:  6,
		7:  11,
		8:  8,
		9:  5,
		10: 10,
		11: 7,
		12: 12,
	}
)

func getCenturyCode(century int) int {
	return centuryCodes[century]
}

func getMonthDoomsday(month int) int {
	return monthDoomsdays[month]
}
