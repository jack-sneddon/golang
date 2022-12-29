package main

import "fmt"

func main() {
	// Declare a variable x and initialize it to 5
	x := 5

	// Declare a pointer p that points to x
	p := &x

	// Print the value of x and the memory address of x
	fmt.Println("x:", x)
	fmt.Println("&x:", &x)

	// Print the value of p and the value it points to
	fmt.Println("p:", p)
	fmt.Println("*p:", *p)

	// Modify the value of x through the pointer p
	*p = 10

	// Print the new value of x
	fmt.Println("x:", x)
}
