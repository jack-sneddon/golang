package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	playDefer()

	// defer is LIFO
	for i := 0; i < 5; i++ {
		defer fmt.Printf("\n%d ", i)
	}

	// stacking defers goes into a LIFO stack
	defer delay(10)
	defer delay(11)
	defer delay(12)

	// clever use of function timing
	defer duration(track("foo"))

}

func playDefer() {
	defer fmt.Println("playDefer() - point 1")
	x := 500
	fmt.Println("playDefer() - point 2 ", x)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func delay(milisecs int) {
	time.Sleep(time.Duration(milisecs) * time.Millisecond)
	log.Printf("\nslept for %dms", milisecs)
}
