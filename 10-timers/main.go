package main

/*
	Time
	Timers
	Tickers
*/
import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// ***time
	fmt.Println("***Time")
	playTime()

	fmt.Println("***Timers")
	playTimer()

	fmt.Println("***Tickers")
	//playTicker()

	fmt.Println("***Sleep")
	randomSleep()

}

func playTime() {

	printTime()

	// delay ms
	delay(100)
	delay(200)

	// clever way to time a function -
	// to-do: need to find the source again
	deferTimer()
}

func delay(milisecs int) {
	time.Sleep(time.Duration(milisecs) * time.Millisecond)
	log.Printf("\ndelay(%d) ms ==> %v", milisecs, time.Now())
}

func randomSleep() {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

/*
clever use of defer function to wrap a timer around a function
*/
func deferTimer() {
	defer duration(track("foo"))
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

// Timers represent a single event in the future.
func playTimer() {
	// how long will it wait - 2 seconds
	/*
		playTimer := time.NewTimer(2 * time.Second)

		// how is it different than time.Sleep?
		// allows me to cancel before it finishes

		<-playTimer.C
		fmt.Println("Timer 1 fired")

		timer2 := time.NewTimer(time.Second)

	*/

	// The <-timer1.C blocks on the timer’s channel C
	// until it sends a value indicating that the timer fired.

}

func playTicker() {
	panic("unimplemented")
}

func printTime() {
	now := time.Now()
	fmt.Println("Current Time: ", now.Format("15:04:05"))
	fmt.Println("Current Date:", now.Format("Jan 2, 2006"))
	fmt.Println("Current Timestamp:", now.Format(time.Stamp))
	fmt.Println("Current ANSIC:", now.Format(time.ANSIC))
	fmt.Println("Current UnixDate:", now.Format(time.UnixDate))
	fmt.Println("Current Kitchen Time:", now.Format(time.Kitchen))
}
