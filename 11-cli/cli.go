package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
e.g.
$ go run *.go preble east west southwest ...
*/
func main() {
	fmt.Println("== CLI ")

	// if you don't provide a command argument, there will be an error
	commandLineArguments()

	readSingleLine()

	label := "Please enter your name: "
	name := promptForName(label)

	fmt.Printf("Hello %s!\n", name)

}

func promptForName(namePrompt string) string {
	fmt.Println("==Prompt for Name")
	var str string

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, namePrompt+" ")
		str, _ = r.ReadString('\n')
		if str != "" {
			break
		}
	}
	return strings.TrimSpace(str)

}
