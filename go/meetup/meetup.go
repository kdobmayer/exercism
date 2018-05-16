package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	startDay, endDay := dayRange(week, month, year)
	for i := startDay; i <= endDay; i++ {
		date := time.Date(year, month, i, 0, 0, 0, 0, time.UTC)
		if date.Weekday() == weekday {
			return i
		}
	}
	return 0
}

func dayRange(week WeekSchedule, month time.Month, year int) (int, int) {
	switch week {
	case First:
		return 1, 7
	case Second:
		return 8, 14
	case Third:
		return 15, 21
	case Fourth:
		return 22, 28
	case Last:
		switch month {
		case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
			return 25, 31
		case time.February:
			if year%4 == 0 {
				return 23, 29
			}
			return 22, 28
		default:
			return 24, 30
		}
	case Teenth:
		return 13, 19
	default:
		return 0, 0
	}
}