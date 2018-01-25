package clock

import "fmt"

const minutesInDay int = 24 * 60

type Clock struct {
	minutes int
}

func (self *Clock) normalize() {
	self.minutes %= minutesInDay
	if self.minutes < 0 {
		self.minutes += minutesInDay
	}
}

func New(hour, minute int) Clock {
	clock := Clock{minute + (hour * 60)}
	clock.normalize()
	return clock
}

func (self Clock) String() string {
	hour := self.minutes / 60
	minute := self.minutes % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

func (self Clock) Add(minutes int) Clock {
	newClock := Clock{self.minutes + minutes}
	newClock.normalize()
	return newClock
}
