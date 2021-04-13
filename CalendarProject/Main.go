package main

import "fmt"

type Date struct {
	day,month,year int
}
func (d Date) toString() string {
	return fmt.Sprintf("%d.%d.%d", d.day, d.month, d.year)
}

func compare(first, second Date) int {

	// returns "-1" if first < second
	// returns "1" if first > second
	// returns "0" if first = second

	if first.year < second.year {
		return 1
	}
	if first.year > second.year {
		return -1
	}

	if first.month < second.month {
		return 1
	}
	if first.month > second.month {
		return -1
	}

	if first.day < second.day {
		return 1
	}
	if first.day > second.day {
		return -1
	}

	return 0
}

func getOldestDate(firstDate, secondDate Date) Date {
	var start Date
	var end Date
	switch compare(firstDate, secondDate) {
	case -1:
		{
			start = firstDate
			end = secondDate
			break
		}
	case 1:
		{
			start = secondDate
			end = firstDate
			break
		}
	case 0:
		{
			return start
		}
	}
	return end
}

func getDateDifference(start, end Date, yearDifference, monthDifference, weekDifference, dayDifference int) {

	var days, weeks, months, years string

	if yearDifference == 1 {
		years = "year"
	} else {
		years = "years"
	}

	if monthDifference == 1 {
		months = "month"
	} else {
		months = "months"
	}

	if weekDifference == 1 {
		weeks = "week"
	} else {
		weeks = "weeks"
	}

	if dayDifference == 1 {
		days = "day"
	} else {
		days = "days"
	}

	fmt.Println("Difference between ", start, "and", end, "is",
		yearDifference, years, monthDifference, months, weekDifference, weeks, "and", dayDifference, days)
}
func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func getDaysInMonth(month, year int) int {
	var days int
	switch month {
	case 1:
		{
			days = 31
			break
		}
	case 2:
		{
			if isLeapYear(year) {
				days = 29
			} else {
				days = 28
			}
			break
		}
	case 3:
		{ days = 31
			break
		}
	case 4:
		{ days = 30
			break
		}
	case 5:
		{ days = 31
			break
		}
	case 6:
		{ days = 30
			break
		}
	case 7:
		{ days = 31
			break
		}
	case 8:
		{ days = 31
		break
		}
	case 9:
		{ days = 30
			break
		}
	case 10:
		{ days = 31
			break
		}
	case 11:
		{ days = 30
			break
		}
	case 12:
		{ days = 31
			break
		}
	default:
		{ days = 31
			break
		}
	}
	return days
}

func getStartAndEndDate(firstDate, secondDate Date) (start, end Date) {
	isDateOlder := false

	if firstDate.year < secondDate.year {
		isDateOlder = true
	}

	if firstDate.year == secondDate.year && firstDate.month < secondDate.month {
		isDateOlder = true
	}

	if firstDate.year == secondDate.year && firstDate.month == secondDate.month && firstDate.day < secondDate.day {
		isDateOlder = true
	}

	var startYear int
	var startMonth int
	var startDay int

	var endYear int
	var endMonth int
	var endDay int

	if isDateOlder {
		startYear = firstDate.year
		endYear = secondDate.year

		startMonth = firstDate.month
		endMonth = secondDate.month

		startDay = firstDate.day
		endDay = secondDate.day
	} else {
		startYear = secondDate.year
		endYear = firstDate.year

		startMonth = secondDate.month
		endMonth = firstDate.month

		startDay = secondDate.day
		endDay = firstDate.day
	}

	return Date{year: startYear, month: startMonth, day: startDay}, Date{year: endYear, month: endMonth, day: endDay}
}

func getDifferenceInDays(firstDate, secondDate Date) int {
	totalDays := 0

	start, end := getStartAndEndDate(firstDate, secondDate)

	currentDay := start.day
	currentMonth := start.month
	currentYear := start.year

	endDay := end.day
	endMonth := end.month
	endYear := end.year

	var currentMonthLimit int

	for {

		if currentDay == endDay && currentMonth == endMonth && currentYear == endYear {
			break
		}

		totalDays++
		currentDay++

		currentMonthLimit = getDaysInMonth(currentMonth, currentYear)

		if currentDay >= currentMonthLimit {
			currentDay = 1
			currentMonth++
			if currentMonth >= 12 {
				currentMonth = 0
				currentYear++
			}
		}
	}
	fmt.Println("Difference between ",start," and ",end, " in days is: ",totalDays)

	return totalDays
}
func addDays(date Date, addedDays int) Date {

	var monthLimit int

	originalDaysToAdd := addedDays

	startDay := date.day
	startMonth := date.month
	startYear := date.year

	currentDay := startDay
	currentMonth := startMonth
	currentYear := startYear

	for {

		if addedDays <= 1 {
			break
		}

		currentDay++

		monthLimit = getDaysInMonth(currentMonth, currentYear)

		if currentDay >= monthLimit {
			currentDay = 1
			currentMonth++
			if currentMonth >= 12 {
				currentMonth = 1
				currentYear++
			}
		}

		addedDays--
	}

	endDate := Date{year: currentYear, month: currentMonth, day: currentDay}
	fmt.Println("Add", originalDaysToAdd, "days to", date, "= ", endDate)

	return endDate
}

func subtractDaysFromDate(date Date, subtractedDays int) Date {

	var monthLimit int

	originalDaysToSubtract := subtractedDays

	startDay := date.day
	startMonth := date.month
	startYear := date.year

	currentDay := startDay
	currentMonth := startMonth
	currentYear := startYear



	for {

		if subtractedDays == 0 {
			break
		}

		currentDay--

		monthLimit = getDaysInMonth(currentMonth, currentYear)

		if currentDay < 1 {
			currentDay = monthLimit
			currentMonth--
			if currentMonth >= 12 {
				currentMonth = 12
				currentYear--
			}
		}

		subtractedDays--
	}

	endDate := Date{year: currentYear, month: currentMonth, day: currentDay}
	fmt.Println("Subtract", originalDaysToSubtract, "days from", date, "= ", endDate)

	return endDate
}


func main() {
	d1 := Date{day: 10, month: 1, year: 1999}
	d2 := Date{day: 18, month: 12, year: 1998}

	fmt.Println(d1)
	fmt.Println(d2)

	fmt.Println("")
	fmt.Println("Comparisons: ")
	fmt.Println("Oldest date between ", d1, " and ", d2, " is ", getOldestDate(d1, d2))

	fmt.Println("")
	days := 20
	fmt.Println("Add / Subtract", days, "days to date:")
	addDays(d1, days)
	addDays(d2, days)
	subtractDaysFromDate(d1, days)
	subtractDaysFromDate(d2, days)
	fmt.Println("--------------------------------------------------------------------------------")
}
