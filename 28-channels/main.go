package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=======concurrency======")

	playChannelUnbuffered()

	playChannelBuffering()

	playClosingChannels()

	playRangeChannels()

	helloWorldOnce()

}

func getTime(millisecs int, c chan string) {
	fmt.Printf("getTime() -  started and waiting > ")
	time.Sleep(time.Duration(millisecs) * time.Millisecond)
	currentTime := time.Now()

	// send timestamp to channel
	c <- currentTime.Format(time.Stamp)
	fmt.Println("getTime() - completed > ")
}

func getTimeStamp(in <-chan int, out chan<- string) {
	fmt.Println("Initializing getTimeStamp subroutine...")
	millisecs := <-in
	fmt.Println("received channel input ...", millisecs)
	time.Sleep(time.Duration(millisecs) * time.Millisecond)
	currentTime := time.Now()

	out <- currentTime.Format(time.Stamp) // send time to out
}

func playChannelUnbuffered() {

	fmt.Println("\n=========\nBuffered Channels...\n========\n")
	fmt.Println("\n***Use an unbuffered channel to get the time after a delay")
	c := make(chan string)

	// receive from channel
	go getTime(2000, c)
	fmt.Println(<-c)

	// watch the sequence of the back and forth
	fmt.Printf("\n\n| main() - start sequence > ")
	go getTime(2000, c)

	time.Sleep(2 * time.Second)
	fmt.Printf(" main() - Back and ready to get channel info! > ")
	secondTime := <-c
	time.Sleep(2 * time.Second)

	fmt.Printf(" main() - Time received from channel... | ")
	fmt.Println(secondTime)

	go getTime(3000, c)
	go getTime(20, c)

	// receive from multiple channels
	x, y := <-c, <-c // receive from c

	fmt.Println("\nDate and time A: ", x)
	fmt.Println("\nDate and time B: ", y)

	fmt.Println("\n***Distribute work by spawning persistent workers and send and receive information through channels.")
	//send and receive data from channel
	in := make(chan int)
	out := make(chan string)

	// spawn persistent workers
	go getTimeStamp(in, out)
	go getTimeStamp(in, out)
	go getTimeStamp(in, out)
	go getTimeStamp(in, out)

	go func() {
		in <- 1000
		in <- 2000
		in <- 1500
		in <- 50
	}()

	// Now we wait for each result to come in
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

}

// chan<- int
// out chan<- int
func playChannelBuffering() {
	fmt.Println("\n=========\nBuffered Channels...\n========")

	fmt.Println("\n***Use a buffered channel to get the time after a delay")
	// notice the second parameter denotes the array size of the buffer - capacity
	c := make(chan string, 2)

	// watch the sequence of the back and forth
	fmt.Printf("\n\n||main() - Start sequence ... > ")
	go getTime(1000, c)

	time.Sleep(2 * time.Second)
	fmt.Printf(" main() - Back and ready to get channel info! >")
	secondTime := <-c
	time.Sleep(2 * time.Second)

	fmt.Println(" main() - Time received from channel... ||")
	fmt.Println(secondTime)

}

func playClosingChannels() {
	fmt.Println("\n=========\nClosing Channels...\n========")
	c := make(chan string, 2)

	// run the function, but close it before we retrieve it
	go getTime(1000, c)
	close(c)

	secondTime, ok := <-c
	if !ok {
		fmt.Println("channel is closed - zero value returned")
	} else {
		fmt.Println("channel is open - value returned")
	}

	fmt.Println(secondTime)
	fmt.Println("*******")

}

func channelWork(ch chan string) {
	// range allows for the channel to keep reading until the channel is closed
	for data := range ch {
		fmt.Println("Channel is open", data)
	}
	fmt.Println("Channel is closed")
}

func playRangeChannels() {
	fmt.Println("\n=========\nRange of Channels...\n========")
	// to-do
	// channelWork()

}

func playSelectChannels() {
	//doWork()
}

func doWork(ch1, ch2 chan string) {
	select {
	case msg1 := <-ch1:
		fmt.Println("received from ch1", msg1)
	case msg2 := <-ch2:
		fmt.Println("received from ch2", msg2)
	default:
		fmt.Println("nothing received")

	}
}

// purpose of the done chanel is to signal it's done
// empty struct to keep memory low as possible
func workDonePattern(input <-chan string, done <-chan struct{}) {

	for {
		select {
		case in := <-input:
			fmt.Println("Got some input: ", in)
		case <-done:
			return
		}
	}
}

func helloWorldOnce() {
	var once sync.Once
	for i := 0; i < 50; i++ {
		once.Do(func() {
			fmt.Println("Hello World")
		})
	}
}
