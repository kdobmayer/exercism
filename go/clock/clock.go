package clock

import (
	"fmt"
)

type Clock int

func New(hour, minute int) Clock {
	minutes := (hour*60 + minute) % 1440
	if minutes < 0 {
		minutes += 1440
	}
	return Clock(minutes)
}

func (c Clock) String() string {
	hour := c / 60 % 24
	minute := c % 60
	return fmt.Sprintf("%.2d:%.2d", hour, minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}