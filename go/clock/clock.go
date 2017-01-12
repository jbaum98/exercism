// Package clock provides tools for working with times.
package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour, minute int
}

// New creates a new Clock, ensuring that the resulting clock has hour ∈ [0,24) and minute ∈ [0,60).
func New(hour, minute int) Clock {
	hour += minute / 60
	minute = minute % 60
	if minute < 0 {
		minute += 60
		hour--
	}
	hour %= 24
	if hour < 0 {
		hour += 24
	}
	return Clock{
		hour:   hour,
		minute: minute,
	}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	return New(clock.hour, clock.minute+minutes)
}
