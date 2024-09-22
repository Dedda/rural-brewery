package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	weekdays = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	months   = [12]string{"Janueary", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
)

type GameCalendar struct {
	year    int
	month   int
	day     int
	hour    int
	minute  int
	counter int
}

func (c *GameCalendar) Update() {
	const (
		ticksPerMinute = 60
		minutesPerHour = 60
		hoursPerDay    = 24
		daysPerMonth   = 28
		monthsPerYear  = 12
	)
	c.counter++
	if c.counter >= ticksPerMinute {
		c.counter -= ticksPerMinute
		c.minute++
		fmt.Println(fmt.Sprintf("%s, %d of %s, %s", weekdays[c.month], c.day+1, months[c.month], c.FormatTimeOfDay()))
	}
	if c.minute >= minutesPerHour {
		c.minute -= minutesPerHour
		c.hour++
	}
	if c.hour >= hoursPerDay {
		c.hour -= hoursPerDay
		c.day++
	}
	if c.day >= daysPerMonth {
		c.day -= daysPerMonth
		c.month++
	}
	if c.month >= monthsPerYear {
		c.month -= monthsPerYear
		c.year++
	}
}

func (c *GameCalendar) FormatTimeOfDay() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c *GameCalendar) NameOfWeekday() string {
	return weekdays[c.day%7]
}

func (c *GameCalendar) NameOfMonth() string {
	return months[c.month%12]
}

func (g *GameCalendar) Draw(screen *ebiten.Image, position Point2F) {

}
