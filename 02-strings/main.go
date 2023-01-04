package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("***String replacement")

	message := "first string value"
	fmt.Println(message)

	message = "second string value"
	fmt.Println(message)

	// check what is happening in memory
	fmt.Println("\n***String replacement 2")
	messagePtr := &message
	message = "third string value"

	fmt.Println("message:", message)
	fmt.Println("*message:", *messagePtr)

	// embedded quotes
	fmt.Println("\n***String embedded quotes and apostrophes")

	message = "this is a \"quote\""
	fmt.Println("message:", message)
	message = "this is another 'quote'"
	fmt.Println("message:", message)

	// replace a word in the string
	fmt.Println("\n***Replace a word in a string")
	message = "I like dogs"
	messagePtr = &message

	fmt.Println("message:", message)

	// creates a new string and assigns it to the variable
	message = strings.Replace(message, "dogs", "guinea pigs", 1)
	fmt.Println("message:", message)

	//show that it creates a new string
	fmt.Println("message loc:", &message)
	fmt.Println("message ptr:", messagePtr)
	fmt.Println("message ptr:", *messagePtr)

	fmt.Println("\n***UPPERCASE / lowercase")
	message = strings.ToUpper(message)
	fmt.Println("message:", message)
	message = strings.ToLower(message)
	fmt.Println("message:", message)

	message = "Indiana Jones and the temple of Doom"
	message = strings.Title(message)
	fmt.Println("message:", message)

	fmt.Println("\n***tabs and newlines")
	message = "this is a \n\tsplit line"
	fmt.Println("message", message)

	fmt.Println("\n***remove whitespace")
	message0 := "@@this is a message@@"
	message1 := "     this is a message     "
	message2 := "     this is a message     "
	message3 := "     this is a message     "
	message4 := "     this is a message     "

	// trim leading/trailing starting at the specified substring
	message0 = strings.Trim(message0, "@@")
	// leading and trailing whitepsace
	message1 = strings.TrimSpace(message1)
	message2 = strings.TrimLeft(message2, " ")
	message3 = strings.TrimRight(message3, " ")
	message4 = strings.ReplaceAll(message4, " ", "")

	fmt.Println("trim whitespace 0>", message0)
	fmt.Println("trim whitespace 1>", message1)
	fmt.Println("trim whitespace 2>", message2)
	fmt.Println("trim whitespace 3>", message3)
	fmt.Println("trim whitespace 4>", message4)

	fmt.Println("\n***treat strings like an array")
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	// Convert the ASCII value to a character
	character := string(alphabet[17])
	fmt.Println("the 14th letter of the alphabet is >", character)

	// string length is fast - it is just returning the value which is part of the stored value
	fmt.Println("\n***string length")
	fmt.Println("lenth of the alphabet is >", len(alphabet))

	// split a string
	fmt.Println("\n***string split")
	name := "Frank Lloyd Wright"
	fmt.Println(name[6:11]) // middle name

	names := strings.Split(name, " ") // create substrings
	fmt.Println(names[1])             // print middle name

	name2 := strings.Join(names, " ")
	fmt.Println(name2) // join names array

	// string literal (notice the lack of \n)
	fmt.Println("\n***string literal")
	poem :=
		`Some say the world will end in fire,
	Some say in ice.
	From what I’ve tasted of desire
	I hold with those who favour fire.
	But if it had to perish twice,
	I think I know enough of hate
	To say that for destruction ice
	Is also great
	And would suffice. -- Robert Frost`

	fmt.Println(poem) // join names array

	// count occurances of a word
	countIce := strings.Count(poem, "ice")
	fmt.Println("Number of times 'ice' is found >", countIce) // join names array

	// first occurence of a word
	firstIce := strings.Index(poem, "ice")
	fmt.Println("first 'ice' is at index >", firstIce) // join names array

	// last occurance of a word
	lastIce := strings.LastIndex(poem, "ice")
	fmt.Println("last 'ice' is at index >", lastIce) // join names array

	// string runes
	playRunes()

	// string comparison
	playStringComparison()

	// substrings
	playSubStrings()
}
