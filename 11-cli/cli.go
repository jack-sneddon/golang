package main

import (
	"fmt"
)

/*
e.g.
$ go run *.go preble east west southwest
*/
func main() {
	fmt.Println("== CLI ")

	commandLineArguments()

	readSingleLine()

}
