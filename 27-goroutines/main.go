package main

import (
	"fmt"
	"sync"
	"time"
)

/*
- Goroutines are functions that are run concurrently
- Goroutines are a way to handle operations concurrently.
- Goroutines is lightweight thread managed by the Go runtime.
- Goroutines run in the same address space, so access to shared memory must be synchronized
*/
func main() {
	playGoRoutines()

	// WARNING - use of sleep is a poor way to manage concurrency
	// because it does not provide guarantees of completion!!!

	// safety with atomic counters
	playWaitGroup()

	// safety with mutex
	playMutex()
}

func playGoRoutines() {
	fmt.Println("======\nPlay Go Routines\n=======")
	go delayRun(3)
	go delayRun(4)
	go delayRun(1)
	go delayRun(5)

	// deliberate delay to wait for above subroutines to complete
	fmt.Println("Sleep in main for 3 seconds...")
	time.Sleep(3 * time.Second)
}

func delayRun(delay time.Duration) {
	delay *= 100000000
	time.Sleep(delay)
	fmt.Println("Time delay = " + delay.String())
}

// Use pointer - need to pass by reference because the counter needs
// to be managed by all the calls.
func delayRunWait(delay time.Duration, wg *sync.WaitGroup) {
	// ensure that done is called by using the 'defer' keyword
	defer wg.Done()

	delayRun(delay)
}

func playWaitGroup() {

	fmt.Println("======\nGo WaitGroup\n=======")

	var wg sync.WaitGroup // wg counter = 0
	// add number of go routines to wait for
	wg.Add(4) // wg counter = 4

	go delayRunWait(2, &wg) //pass by reference, so wg counter is actually zero
	go delayRunWait(3, &wg)
	go delayRunWait(6, &wg)
	go delayRunWait(1, &wg)

	// block until all go routines have completed (or more precicely the inner wait counter reaches zero)
	wg.Wait()

}

type SharedCount struct {
	mu    sync.Mutex
	count int
}

func newSharedCount(start int) *SharedCount {
	sc := SharedCount{}
	sc.count = start
	return &sc
}

func (sc *SharedCount) increaseSharedCount(increase int, wg *sync.WaitGroup) {
	// lock this function
	sc.mu.Lock()
	// make sure we unlock before we continue
	defer sc.mu.Unlock()

	sc.count += increase
	fmt.Printf("Increasing count to %d\n", sc.count)

	if wg != nil {
		wg.Done()
	}

}

func playMutex() {
	fmt.Println("======\nPlay Sync.Mutex (lock/unlock)\n=======")

	sc := newSharedCount(0)

	var wg sync.WaitGroup
	wg.Add(3)

	go sc.increaseSharedCount(300, &wg)
	go sc.increaseSharedCount(200, &wg)
	go sc.increaseSharedCount(100, &wg)

	wg.Wait()

	fmt.Println("Shared Count = ", sc.count)
}
