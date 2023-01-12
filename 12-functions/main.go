package main

import "fmt"

/*
- Go is pass by value by default.
- Go is does allow pass by value by pointers
*/
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
	passReference := 8
	fmt.Printf("\npassReference %d, %p", passReference, &passReference)
	passByReference(&passReference)
	fmt.Printf("\npassReference %d, %p\n", passReference, &passReference)

	// multiple return values
	analysis, status := doSomeCoolThing()
	fmt.Printf("\nAnalysis = %v : status code = %d", analysis, status)

	// Variadic functions - open number of arguments
	fmt.Println("\n")
	printStudents("Joe", "Jane", "Sally", "Bob")

	/* Arrays in Go are values: when you pass an array to a function, it gets a copy of the original array data.
	if you wan ta  function to update its elewments, use a slice that refers to the array.
	*/

}

func doSomeCoolThing() (string, int) {
	status := 1000
	analysis := "it worked!"

	return analysis, status
}

func printStudents(students ...string) {
	fmt.Print(students, " ")
	/*
		for i , student : range students {
			fmt.Println(student)
		}
	*/

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
	fmt.Printf("\npassedReference %d, %p", *passedReference, passedReference)
	*passedReference += 5
	fmt.Printf("\npassedReference += 5 = %v, %p", *passedReference, passedReference)
}
