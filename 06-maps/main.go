package main

import "fmt"

func main() {

	// simple map
	myMap := map[string]int{"biz": 1, "baz": 2, "bar": 3}
	fmt.Println("map:", myMap)

	// make a map of airport codes

	// make(map[key-type]val-type).
	airportCodes := make(map[string]string)

	airportCodes["MKE"] = "Milwaukee"
	airportCodes["JLN"] = "Joplin"
	airportCodes["SLC"] = "Salt Lake City"
	airportCodes["MCI"] = "Kansas City"
	airportCodes["BOS"] = "Boston"
	airportCodes["TUS"] = "Tuscon"
	airportCodes["RDU"] = "Raleigh"
	airportCodes["BLI"] = "Bellingham"
	airportCodes["ORD"] = "Chicago O'Hare"
	airportCodes["SGF"] = "Springfild"

	fmt.Println("\nnum of cities:", len(airportCodes))
	fmt.Println("Airport Codes:", airportCodes)

	// get the value
	city := airportCodes["BLI"]
	fmt.Println("BLI: ", city)

	// remove an entry
	delete(airportCodes, "SGF")
	fmt.Println("\nnum of cities:", len(airportCodes))
	fmt.Println("airport Codes:", airportCodes)

	// look for the presence of a key (prs)
	city, codeExists := airportCodes["SLC"]
	fmt.Println("airport code exist? : ", codeExists, city)

	// ignore the city - we don't need it
	_, codeExists = airportCodes["BUF"]
	fmt.Println("airport code exist? : ", codeExists)

}
