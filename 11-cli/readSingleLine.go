package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readSingleLine() {

	fmt.Println("==Read Command Line")

	readSingleLineWithFmt()

	readSingleLineWithBufioReader()

	readSingleLineWithBufioScanner()

}

/*
Use fmt.Scanln() when you want to read each word individually
*/
func readSingleLineWithFmt() {
	// use fmt.Scanln

	var firstName string
	var lastName string
	var age int

	fmt.Println("\n-\nEnter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Print("\nEnter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Println("\nEnter your age: ")
	fmt.Scanln(&age)

	fmt.Printf("\nHello %s %s.  You are wise for your %d years!\n", firstName, lastName, age)

	/*
		fmt.Println("fmt.Scanln approach -- input exactly three strings:")
		var input1, input2, input3 string
		n, err := fmt.Scanln(&input1, &input2, &input3)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nNumber of items read: %d\n", n)
		fmt.Printf("\nRead line: %s %s %s-\n", w1, w2, w3)
	*/

}

/*
Use Bufio.Reader() when you want to read the entire line including the trailing newline character
*/
func readSingleLineWithBufioReader() {
	fmt.Print("\n-\nEnter your name: ")

	reader := bufio.NewReader(os.Stdin)

	fullName, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your name is: %s\n", fullName)
}

/*
Use Bufio.Reader() when you want to read the entire line without the newline character
*/
func readSingleLineWithBufioScanner() {
	fmt.Println("\n-\nEnter your name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Welcome : %s\n", scanner.Text())
}
