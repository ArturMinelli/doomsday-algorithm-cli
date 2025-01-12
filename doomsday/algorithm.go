package doomsday

import (
	"time"
)

type Variables struct {
	Century                  int
	Decade                   int
	MonthDoomsday            int
	DayToMonthDoomsdayOffset int
	DivisionDecadeByTwelve   int
	RemainderDecadeByTwelve  int
	DivisionRemainderByFour  int
	CenturyCode              int
}

type Doomsday struct {
	Date    time.Time
	Weekday int
	Variables
}

func Run(date time.Time) Doomsday {
	year := int(date.Year())
	month := int(date.Month())
	day := int(date.Day())

	century := (year / 100) * 100
	decade := year - century

	centuryCode := getCenturyCode(century)
	monthDoomsday := getMonthDoomsday(month)

	if checkLeapYearException(year, month) {
		monthDoomsday += 1
	}

	dayToMonthDoomsdayOffset := day - monthDoomsday
	divisionDecadeByTwelve := decade / 12
	remainderDecadeByTwelve := decade % 12
	divisionRemainderByFour := remainderDecadeByTwelve / 4

	weekday := (dayToMonthDoomsdayOffset + divisionDecadeByTwelve + remainderDecadeByTwelve + divisionRemainderByFour + centuryCode) % 7

	if weekday < 0 {
		weekday += 7
	}

	return Doomsday{
		Weekday: weekday,
		Date:    date,
		Variables: Variables{
			Century:                  century,
			Decade:                   decade,
			MonthDoomsday:            monthDoomsday,
			DayToMonthDoomsdayOffset: dayToMonthDoomsdayOffset,
			DivisionDecadeByTwelve:   divisionDecadeByTwelve,
			RemainderDecadeByTwelve:  remainderDecadeByTwelve,
			DivisionRemainderByFour:  divisionRemainderByFour,
			CenturyCode:              centuryCode,
		},
	}
}

func checkLeapYearException(year int, month int) bool {
	if year%100 == 0 && year%400 != 0 {
		return false
	}

	return year%4 == 0 && (month == 1 || month == 2)
}
