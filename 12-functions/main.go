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

	// call a function with multiple return values
	nfl, nba, mlb := favoriteProTeams()
	fmt.Printf("\nFavorite NFL team: %v", nfl)
	fmt.Printf("\nFavorite NBA team: %v", nba)
	fmt.Printf("\nFavorite MLB team: %v\n", mlb)

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
	fmt.Println("\n***Variadic Function")
	printStudents("Joe", "Jane", "Sally", "Bob")

	/* Arrays in Go are values: when you pass an array to a function, it gets a copy of the original array data.
	if you wan ta  function to update its elewments, use a slice that refers to the array.
	*/

	// anonymous function
	// short in scope - less than package
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	// call anonmoys functions
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

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

// return multiple values
func favoriteProTeams() (football string, basketball string, baseball string) {
	football = "Green Bay Packers"
	basketball = "Milwaukee Bucks"
	baseball = "Milwaukee Brewers"

	return

}

// note the naming of the return value (passedVal.  we could have just said "int")
func passByValue(passedValue int) {
	fmt.Printf("\npassedValue %d, %p", passedValue, &passedValue)
	passedValue += 5
	fmt.Printf("\npassedValue += 5 = %v, %p", passedValue, &passedValue)
	// note that passedValue is the return value as defined in the function interface
	return
}

func passByReference(passedReference *int) {
	fmt.Printf("\npassedReference %d, %p", *passedReference, passedReference)
	*passedReference += 5
	fmt.Printf("\npassedReference += 5 = %v, %p", *passedReference, passedReference)
}
