package clock

import "fmt"

const MinIn1Hour = 60
const MinIn1Day = 1440

type Clock struct {
	min int
}

func New(hour, minute int) Clock {
	return Clock{0}.Add(MinIn1Hour*hour + minute)
}

func Time(hour, minute int) Clock {
	c := Clock{(MinIn1Hour*hour + minute) % MinIn1Day}
	if c.min < 0 {
		c.min += MinIn1Day
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/MinIn1Hour, c.min%MinIn1Hour)
}

func (c Clock) Add(min int) Clock {
	c.min = (c.min + min) % MinIn1Day
	if c.min < 0 {
		c.min += MinIn1Day
	}
	return c
}
