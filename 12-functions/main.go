package main

import "fmt"

func main() {
	// call a function
	favoriteNBATeam := favNBATeam()
	fmt.Println(favoriteNBATeam)

	// call a function in a local package
	favoriteNFLTeam := favoriteNFLTeam()
	fmt.Println(favoriteNFLTeam)

	// pass by value
	passValue := 3
	fmt.Printf("\npassValue %v, %p", passValue, &passValue)
	passByValue(passValue)
	fmt.Printf("\npassValue %v, %p\n", passValue, &passValue)

	// pass by reference

	/* Arrays in Go are values: when you pass an array to a function, it gets a copy of the original array data.
	if you wan ta  function to update its elewments, use a slice that refers to the array.
	*/

	passReference := 8
	fmt.Printf("\npassReference %d, %p", passReference, &passReference)
	passByReference(&passReference)
	fmt.Printf("\npassReference %d, %p\n", passReference, &passReference)

	make()
}

func favNBATeam() string {
	return "Milwaukee Bucks"
}

func passByValue(passedValue int) {

	fmt.Printf("\npassedValue %d, %p", passedValue, &passedValue)
	passedValue += 5
	fmt.Printf("\npassedValue += 5 = %v, %p", passedValue, &passedValue)
}

func passByReference(passedReference *int) {
	//fmt.Printf("Before AddPtr, Memory Location: %p, Value: %d\n", x, *x)
	fmt.Printf("\npassedReference %d, %p", *passedReference, passedReference)
	*passedReference += 5
	fmt.Printf("\npassedReference += 5 = %v, %p", *passedReference, passedReference)
}
