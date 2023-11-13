package main

import "fmt"

func main() {
	fmt.Println("=====\nConditions\n=====")

	trueVal := true
	falseVal := false

	// IF/ELSE
	if trueVal {
		fmt.Println("IF/Else = true")
	} else {
		fmt.Println("IF/Else = false")
	}

	// ELSE IF - None
	// https://en.wikipedia.org/wiki/Ternary_conditional_operator#Go

	// AND
	if trueVal && falseVal {
		fmt.Println("\nIF x && y = true")
	} else {
		fmt.Println("\nIF x && y = false")
	}

	// OR
	if trueVal || falseVal {
		fmt.Println("\nIF x || y = true")
	} else {
		fmt.Println("\nIF x || y = false")
	}

	// NOT
	if !falseVal {
		fmt.Println("\nIF !x = true")
	} else {
		fmt.Println("\nIF !x = false")
	}

	// NESTED EXPRESSIONS
	if !(falseVal) && (trueVal || falseVal) {
		fmt.Println("\nnested expressions = true")
	} else {
		fmt.Println("\nnested expressions = false")
	}

	// SWITCH STATEMENT
	region := "U.S"

	var cities []string

	cities = append(cities, "Milwaukee")
	cities = append(cities, "Joplin")
	cities = append(cities, "Salt Lake City")
	cities = append(cities, "Kansas City")
	cities = append(cities, "Boston")
	cities = append(cities, "Tuscon")
	cities = append(cities, "Raleigh")
	cities = append(cities, "Bellingham")
	cities = append(cities, "Grand Rapids")
	cities = append(cities, "Springfild")

	for _, city := range cities {
		// fmt.Printf("\ncity: [%d] = %v", i, city)
		region = getRegion(city)
		fmt.Printf("\n%v is in the %v", city, region)
	}

	// swtich on the type
	//t := 100

}

func getRegion(city string) string {
	region := "U.S."

	mke := "Milwaukee"
	slc := "Salt Lake City"
	mci := "Kansas City"
	bos := "Boston"
	tus := "Tuscon"
	rdu := "Raleigh"
	bli := "Bellingham"
	ord := "Chicago"
	spr := "Springfild"

	east := "East"
	south := "South"
	midwest := "Midwest"
	mountainwest := "Mountain West"
	soutwest := "Southwest"
	westCoast := "West Coast"

	switch city {
	case mke, mci, ord:
		region = midwest
	case slc:
		region = mountainwest
	case bos:
		region = east
	case tus:
		region = soutwest
	case rdu:
		region = south
	case bli:
		region = westCoast
	case spr:
		region = "U.S."
	default:
		region = "unknown"
	}

	return region
}
