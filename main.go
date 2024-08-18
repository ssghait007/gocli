package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Define command-line flags
	makedayscount := flag.Bool("makedayscount", false, "Display the difference in years")
	flag.Parse()

	// Define the start date
	dateOfJoining := os.Getenv("DATE_OF_JOINING")
	startDate, _ := time.Parse("2-1-2006", dateOfJoining)
	currentDate := time.Now()

	// Calculate the difference
	yearDiff := currentDate.Year() - startDate.Year()
	monthDiff := int(currentDate.Month()) - int(startDate.Month())
	dayDiff := currentDate.Day() - startDate.Day()

	// Adjust for negative differences
	if dayDiff < 0 {
		monthDiff--
		// Get the previous month
		prevMonth := currentDate.AddDate(0, -1, 0)
		dayDiff += daysInMonth(prevMonth.Year(), int(prevMonth.Month()))
	}

	if monthDiff < 0 {
		yearDiff--
		monthDiff += 12
	}

	// Display the difference
	if *makedayscount {
		fmt.Printf("You are in IT since: \033[32m%d years, \033[33m%d months, \033[34mand %d days.\033[0m\n", yearDiff, monthDiff, dayDiff)
	}
}

// daysInMonth returns the number of days in a given month and year
func daysInMonth(year int, month int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}
