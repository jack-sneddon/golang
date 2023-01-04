package main

import (
	"fmt"
	"reflect"
)

func playRunes() {
	// runes - a data type that stores codes that represent Unicode characters.
	// go does not have a char type, so anything coming out of a string is  automatically typecast to int32
	fmt.Println("\n***rune")
	emoji := '🦬'
	// fmt.Printf("Data type of %c is %T and the rune value is %U \n", emoji, emoji, emoji)
	fmt.Printf("%c - %s \n with a rune value of %U \n", emoji, reflect.TypeOf(emoji), emoji)

	str := "Jack"
	r_array := []rune(str)
	fmt.Printf("Array of rune values for '%v' is %U\n", str, r_array)

}
