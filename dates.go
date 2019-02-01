// Simple date handling
package main

import "fmt"

// simple struct and functions for date only: time is irrelevant
// days/months start from 1, not 0
type Date struct {
	year  int
	month int
	day   int
}

// display a date in ISO whateveritis format
func show(d Date) string {
	return fmt.Sprintf("%d-%02d-%02d", d.year, d.month, d.day)
}

type Period Date

var monthLengths = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func monthLength(month, year int) int {
	if month == 2 && isLeap(year) {
		return 29
	}
	return monthLengths[month-1]
}

func isLeap(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}

func dateRange(start, end Date, interval Period) []Date {
	dates := make([]Date, 0)
	for d := start; isNotAfter(d, end); d = postpone(d, interval) {
		dates = append(dates, d)
	}
	return dates
}

// check if date a hasn't passed b (i.e. a <= b)
func isNotAfter(a, b Date) bool {
	// switch b.year - a.year {
	switch {
	case b.year > a.year:
		return true
	case b.year < a.year:
		return false
	}
	switch {
	case b.month > a.month:
		return true
	case b.month < a.month:
		return false
	}
	return b.day >= a.day
}

// advance a date by a certain interval of days/months/years
// Behaviour is arbitrary if days and months/years given to advance:
// this function advances years, then months, then days
//   e.g. advancing 1mo31day from 15th Feb -> 15th Mar -> 15th Apr
func postpone(date Date, interval Period) Date {

	date.year = date.year + interval.year
	date.month = date.month + interval.month
	date.day = date.day + interval.day

	// if month or date number is higher than it should be
	for {
		// overflow months
		if date.month > 12 {
			date.year += date.month / 12
			date.month = ((date.month - 1) % 12) + 1
		}
		// overflow days -- month by month in the for loop
		currentMonthDays := monthLength(date.month, date.year)
		if date.day <= currentMonthDays {
			break
		}
		date.day -= currentMonthDays
		date.month += 1
	}

	return date
}