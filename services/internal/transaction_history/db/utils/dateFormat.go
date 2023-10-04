package utils

import (
	"time"
)

func areDatesWithinOneMonth(date1 time.Time, date2 time.Time) bool {
	// Calculate the difference in months between the two dates
	year1, month1, _ := date1.Date()
	year2, month2, _ := date2.Date()
	diffYears := year2 - year1
	diffMonths := int(month2) - int(month1) + (diffYears * 12)

	// Check if the difference is less than or equal to 1
	return diffMonths <= 1
}

func areDatesWithinOneWeek(date1, date2 time.Time) bool {
	// Calculate the difference in days between the two dates
	diff := date2.Sub(date1)
	// Calculate the difference in weeks by dividing the difference in days by 7
	diffWeeks := int(diff.Hours() / 24 / 7)
	// Check if the difference is less than or equal to 1
	return diffWeeks <= 1
}

func areDatesWithinOneYear(date1, date2 time.Time) bool {
	// Calculate the difference in years between the two dates
	year1 := date1.Year()
	year2 := date2.Year()
	diffYears := year2 - year1

	// Check if the difference is less than or equal to 1
	return diffYears <= 1
}

func subtractWeeks(date time.Time, weeks int) time.Time {
	// Subtract the specified number of weeks.
	result := date.AddDate(0, 0, -7*weeks) // 7 days in a week
	return result
}

func LookBackView(date string, recurring string) string {
	var newDate string
	inputDate, _ := time.Parse("2006-01-02", date)
	if recurring == "weekly" {
		const threeWeeks = 3
		newDate = subtractWeeks(inputDate, threeWeeks).Format("2006-01-02")
	} else if recurring == "monthly" {
		const threeMonths = 12
		newDate = subtractWeeks(inputDate, threeMonths).Format("2006-01-02")
	} else if recurring == "yearly" {
		const threeYears = 156
		newDate = subtractWeeks(inputDate, threeYears).Format("2006-01-02")
	}
	return newDate
}

func IsTrendHealthy(date1 string, date2 string, view string) bool {
	parseDate1, _ := time.Parse("2006-01-02", date1)
	parseDate2, _ := time.Parse("2006-01-02", date2)

	switch view {
		case "weekly": {
			return areDatesWithinOneWeek(parseDate1, parseDate2)
		}
		case "monthly": {
			return areDatesWithinOneMonth(parseDate1, parseDate2)
		}
		case "yearly": {
			return areDatesWithinOneYear(parseDate1, parseDate2)
		}
		default: {
			return false
		}
	}
}

func PredictNextDueDate(previousDateCycle, lastDateCycle string) string {
	last, _ := time.Parse("2006-01-02", lastDateCycle)
	prev, _ := time.Parse("2006-01-02", previousDateCycle)
	// Calculate the time interval between the two previous due dates
	interval := last.Sub(prev)
	// Add the interval to the most recent due date to predict the next due date
	nextDueDate := last.Add(interval)
	return nextDueDate.Format("2006-01-02")
}
// Test funcitons
