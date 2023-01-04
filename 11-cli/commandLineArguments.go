package main

import (
	"fmt"
	"os"
)

func commandLineArguments() {
	fmt.Println("==Read Command Line")

	// read command line arguments
	argumentsWithProgram := os.Args
	argumentsOnly := os.Args[1:]

	specificArgument := os.Args[1]

	fmt.Printf("\nArguments including the application: '%v'\n", argumentsWithProgram)
	fmt.Printf("\nList of arguments: '%v'\n", argumentsOnly)
	fmt.Printf("\nFirst argument '%v'\n", specificArgument)

}
