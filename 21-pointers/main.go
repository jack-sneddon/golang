package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declare a variable x and initialize it to 5
	x := 5

	// '*' is a pointer.
	var p1 *int
	p1 = &x

	// '&' is the address.  this is an implicit declaration
	p2 := &x

	// Print the value of x and the memory address of x
	fmt.Println("&x:", &x)
	fmt.Println("x:", x)
	fmt.Printf("type of x = %T\n", x)

	// Print the value of p and the value it points to
	fmt.Println("\n&p1:", &p1)
	fmt.Println("p1:", p1)
	fmt.Println("*p1:", *p1)
	fmt.Println("type of p1 = ", reflect.TypeOf(p1))

	// Print the value of p and the value it points to
	fmt.Println("\n&p2:", &p2)
	fmt.Println("p2:", p2)
	fmt.Println("*p2:", *p2)
	fmt.Println("type of p2 = ", reflect.ValueOf(p2).Kind())

	// Modify the value of x through the pointer p
	*p1 = 10

	// Print the new value of x
	fmt.Println("\n new value of x:\n", x)

	// nil pointer issues resulting in:
	// panic: runtime error: invalid memory address or nil pointer dereference

	// trying to use p3 that is nill - best to add a check around it
	var p3 *int32
	if p3 != nil {
		fmt.Println(*p3)
	} else {
		fmt.Println("p3 is nil!")
	}
}
