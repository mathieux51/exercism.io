package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

const minutesPerDay = 24 * 60

func New(hour, minute int) Clock {
	var m int = (60*hour + minute) % minutesPerDay

	if m < 0 {
		m += minutesPerDay
	}

	return Clock{minute: m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, c.minute+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minute-minutes)
}
