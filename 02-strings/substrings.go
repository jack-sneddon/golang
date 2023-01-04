package main

import (
	"fmt"
	"regexp"
	"strings"
)

func playSubStrings() {
	state := "Missouri"
	cityState := "Joplin, Missouri"
	city1 := "Joplin"
	city2 := "Bentonville"

	wisconsinCities := "Green Bay, Appleton, Milwaukee, Waukesha, Oconomowoc, Sheboygan, Neenah, Wauwatosa, Kewaunee, Menasha, Wautoma, Manitowoc, Ashwaubenon, Wausau, Kenosha, Shawano, Oshkosh, Kaukauna, Suamico"

	fmt.Printf("\nDoes '%s' contain '%s' : %t", cityState, city1, strings.Contains(cityState, city1))
	fmt.Printf("\nDoes '%s' contain '%s' : %t", cityState, state, strings.Contains(cityState, state))
	fmt.Printf("\nDoes '%s' contain '%s' : %t", cityState, city2, strings.Contains(cityState, city2))

	// split strings
	fmt.Printf("\n%v", wisconsinCities)
	cities := strings.Split(wisconsinCities, ", ")
	fmt.Println(wisconsinCities)
	fmt.Printf("\nstrings.Split(): %#v\n", cities)
	fmt.Printf("1st, 3rd & 5th cities are: %v, %v, %v", cities[0], cities[2], cities[4])

	// regex split
	regex := regexp.MustCompile(",\\s*")
	cities2 := regex.Split(wisconsinCities, -1)

	fmt.Printf("\nregex strings.Split(): %#v\n", cities2)

}
