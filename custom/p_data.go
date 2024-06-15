package functions

import (
	"errors"
	"time"
)

type Date struct {
	day   int
	month int
	year  int
}

func (d *Date) SetDat(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("ERROR: Incorrect day value")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(m int) error {
	if (m < 1) || (m > 31) {
		return errors.New("ERROR: Incorrect month value")
	}
	d.month = m
	return nil
}

func (d *Date) SetYear(y int) error {
	if (y < 1875) || (y > time.Now().Year()) {
		return errors.New("ERROR: Incorrect year value")
	}
	d.year = y
	return nil
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}
