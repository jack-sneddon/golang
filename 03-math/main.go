package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	/*
		print numbers
	*/
	fmt.Println("***Convert to String")
	myNumber := 125

	str1 := strconv.FormatInt(int64(myNumber), 10)
	str2 := strconv.Itoa(myNumber)
	fmt.Printf("%v, %v\n", str1, str2)

	fmt.Println("***Integer")

	myInt1 := 10
	myInt2 := 2

	myAdd := myInt1 + myInt2
	mySub := myInt1 - myInt2
	myMult := myInt1 * myInt2
	myDiv := myInt1 / myInt2
	// myIntTruncate := 135 //33
	myMod := myInt2 % myInt1

	fmt.Printf("%v + %v = %v\n", myInt1, myInt2, myAdd)
	fmt.Printf("%v - %v = %v\n", myInt1, myInt2, mySub)
	fmt.Printf("%v * %v = %v\n", myInt1, myInt2, myMult)
	fmt.Printf("%v / %v = %v\n", myInt1, myInt2, myDiv)
	//fmt.Printf("%v + %v = %v", myInt1, myInt2, myAdd)
	fmt.Printf("%v mod %v = %v\n", myInt1, myInt2, myMod)

	evenOddNumber := 333
	if evenOddNumber == 0 {
		fmt.Printf("\n%v is neither even or odd\n", evenOddNumber)
	} else {
		if isEven(evenOddNumber) {
			fmt.Printf("\n%v is an even number\n", evenOddNumber)
		} else {
			fmt.Printf("\n%v is an odd number\n", evenOddNumber)
		}
		if isOdd(evenOddNumber) {
			fmt.Printf("\n%v is an odd number\n", evenOddNumber)
		} else {
			fmt.Printf("\n%v is an even number\n", evenOddNumber)
		}
	}

	fmt.Println("***Float")
	pi := math.Pi
	fmt.Printf("compact representation: %v%g\n", pi, pi)
	fmt.Printf("presision: \n %.2f \n(%6.2f)\n", pi, pi)
	fmt.Printf("exponential notation: %e\n", pi)
	fmt.Printf("round to int: %v\n", math.Round(pi))

}

func isEven(checkNum int) bool {
	isEven := false
	if checkNum == 0 {
		isEven = false
	} else if checkNum%2 == 0 {
		isEven = true
	} else {
		isEven = false
	}
	return isEven
}

func isOdd(checkNum int) bool {
	isOdd := false
	if checkNum == 0 {
		isOdd = false
	} else if checkNum%2 == 0 {
		isOdd = false
	} else {
		isOdd = true
	}
	return isOdd
}
