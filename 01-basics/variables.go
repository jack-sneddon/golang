package main

import (
	"fmt"
	"math"
)

func playVariables() {
	fmt.Println("\n===Tour of variables===")

	// quick runthrough of primative types.
	var b bool = true
	var i int = 42
	var f float32 = 3.14
	var c complex64 = 1 + 2i
	var s string = "this is a string"

	// Print the values of the variables
	fmt.Println("\n***Primative Types examples")
	fmt.Println("boolean: ", b)
	fmt.Println("integer: ", i)
	fmt.Println("float32: ", f)
	fmt.Println("complex64: ", c)
	fmt.Println("string: ", s)

	// global assignment
	fmt.Println(ga, gb)

	// constant
	const NOCHANGE string = "Cannot Change"
	fmt.Println(NOCHANGE)

	// zero value - variables start out with zero
	var x int
	var y string
	var z float64
	fmt.Printf("zero init: %v, %v, %v", x, y, z)

	// *** string
	fmt.Println("\n***String")
	var myStr1 string = "My String 1"
	// create a string with shorthand (let compiler figure out the type)
	myStr2 := "My String 2"

	fmt.Println("myNumber = " + myStr1)
	fmt.Println("myNumber2 = " + myStr2)

	// *** iteger
	fmt.Println("\n***Integer")
	var int1, int2 int = 1, 2
	fmt.Println("int1, int2", int1, int2)
	fmt.Println("int1:", int1)
	fmt.Printf("%v\n", int1)
	fmt.Printf("%d\n", int2)

	// *** boolean
	fmt.Println("\n\n***Boolean")
	var bool1 = true
	fmt.Printf("%v %t\n", bool1, bool1)

	// *** float
	fmt.Println("\n***Float")

	pi := math.Pi
	fmt.Printf("\ncompact representation: %v%g", pi, pi)
	fmt.Printf("\npresision: \n %.2f \n(%6.2f)", pi, pi)
	fmt.Printf("\nexponential notation: %e", pi)

	// *** print the type
	fmt.Println("\n\n***Types")
	fmt.Printf("%T %T\n", myStr1, &myStr1)
	fmt.Printf("%T %T\n", int1, &int2)
	fmt.Printf("%T %T\n", bool1, &bool1)
	fmt.Printf("%T %T\n", pi, &pi)

}
