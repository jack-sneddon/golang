package main

import (
	"fmt"
	"strings"
)

func playStringComparison() {

	fmt.Println("\n===String Comparisons")

	string1 := "Utah"
	string2 := "Wisconsin"
	/*
		string3 := "Missouri"
		string4 := "North Salt Lake, Utah"
		string5 := "Green Bay, Wisconsin"
		string6 := "Milwaukee, Wisconsin"
		string7 := "Joplin, Missouri"
	*/
	string8 := "Utah"
	string9 := "UTAH"

	fmt.Println("'=='")
	fmt.Printf("%s == %s: %t (%p - %p)\n", string1, string2, string1 == string2, &string1, &string2)
	fmt.Printf("%s == %s: %t (%p - %p)\n", string1, string8, string1 == string8, &string1, &string2)

	fmt.Println("'!='")
	fmt.Printf("%s != %s: %t (%p - %p)\n", string1, string2, string1 != string2, &string1, &string2)
	fmt.Printf("%s != %s: %t (%p - %p)\n", string1, string8, string1 != string8, &string1, &string2)

	fmt.Println("'>='")
	fmt.Printf("%s >= %s: %t (%p - %p)\n", string1, string2, string1 >= string2, &string1, &string2)
	fmt.Printf("%s >= %s: %t (%p - %p)\n", string1, string8, string1 >= string8, &string1, &string2)

	fmt.Println("'<='")
	fmt.Printf("%s <= %s: %t (%p - %p)\n", string1, string2, string1 <= string2, &string1, &string2)
	fmt.Printf("%s <= %s: %t (%p - %p)\n", string1, string8, string1 <= string8, &string1, &string2)

	// compare function
	fmt.Println("Compare Function")
	fmt.Printf("strings.Compare(%s, %s): %d\n", string1, string2, strings.Compare(string1, string2))
	fmt.Printf("strings.Compare(%s, %s): %d\n", string1, string8, strings.Compare(string1, string8))
	fmt.Printf("strings.Compare(%s, %s): %d\n", string2, string1, strings.Compare(string2, string1))

	// case insensive comparison
	fmt.Printf("strings.Compare(%s, %s): %d\n", string1, string9, strings.Compare(string1, string9))
	fmt.Printf("strings.EqualFolds(%s, %s): %v\n", string1, string9, strings.EqualFold(string1, string8))

}
