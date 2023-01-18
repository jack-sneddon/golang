package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\n===== Date and Time ====")

	playDateAndTime()

	playTimers()

	playTickers()

}

// Timers are for when you want to do something once in the future
func playTimers() {
	fmt.Println("\n***Timers")

}

// Tickers are for when you want to do something repeatedly at regular intervals.
func playTickers() {
	fmt.Println("\n***Tickers")

}

func playDateAndTime() {
	fmt.Println("\n***Date and Time")
	currentTime := time.Now()

	fmt.Printf("\nCurrent date and time = %v", currentTime.String())
	fmt.Printf("\nMM-DD-YYYY : %s", currentTime.Format("01-02-2006"))
	fmt.Printf("\nYYYY-MM-DD : %s", currentTime.Format("2006-01-02"))
	fmt.Printf("\nYYYY.MM.DD : %s", currentTime.Format("2006.01.02 15:04:05"))
	fmt.Printf("\nYYYY-MM-DD hh:mm:ss : %s", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("\nYYYY/MM/DD hh:mm:ss : %s", currentTime.Format("2006/01/02 15:04:05"))
	fmt.Printf("\nTime with MicroSeconds: %s", currentTime.Format("2006-01-02 15:04:05.000000"))
	fmt.Printf("\nTime with NanoSeconds: %s", currentTime.Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("\nMonth Number : %s", currentTime.Format("2006-1-02"))
	fmt.Printf("\nMonth Name : %s", currentTime.Format("2006-January-02"))
	fmt.Printf("\nMonth Abbreviation : %s", currentTime.Format("2006-Jan-02"))

	// get the current year, month, day only
	fmt.Printf("\nCurrent Year  : %v", currentTime.Year())
	fmt.Printf("\nCurrent Month  : %v", currentTime.Month())
	fmt.Printf("\nCurrent Day  : %v", currentTime.Day())
	fmt.Printf("\nDay of Weekday  : %v", currentTime.Weekday())
	fmt.Printf("\nDay of Hour  : %v", currentTime.Hour())
	fmt.Printf("\nDay of Year  : %v", currentTime.YearDay())
	fmt.Printf("\nTimezone  : %v", currentTime.Location())

}
