package main

import (
	"fmt"
	"time"
)

func calculateEasterSunday(year int) time.Time {
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451
	month := (h + l - 7*m + 114) / 31
	day := ((h + l - 7*m + 114) % 31) + 1

	// Convert to Julian calendar if date is before September 14, 1752
	if year < 1752 || (year == 1752 && month < 9) || (year == 1752 && month == 9 && day < 14) {
		return time.Date(int(year), time.Month(month), int(day-10), 0, 0, 0, 0, time.UTC)
	}

	return time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
}

func main() {
	var yearInput int

	fmt.Print("Enter the year (either two digits or four digits):\n> ")
	_, err := fmt.Scanf("%d", &yearInput)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// If the user enters a two-digit year, convert it to a four-digit year
	if yearInput == 0 || yearInput == 00{
		yearInput = 2000
	}
	
	if yearInput >= 1 && yearInput < 100 {
		currentYear := time.Now().Year()
		century := currentYear / 100 * 100
		yearInput += century

		// If the resulting year is in the past, add 100 years
		if yearInput < currentYear {
			yearInput += 100
		}
	}

	easterSunday := calculateEasterSunday(yearInput)
	fmt.Printf("Easter Sunday in %d is on %s\n", yearInput, easterSunday.Format("02/01/2006"))
}
