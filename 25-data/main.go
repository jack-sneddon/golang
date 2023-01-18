package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("\n======= \nData\n======= ")

	createCvsFile()

}

func createCvsFile() {
	file, err := os.OpenFile("data/test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}

	heading := []string{"Band", "Song", "Rating"}
	band1 := []string{"Gorillaz", "Feel Good Inc", "5"}
	band2 := []string{"Van Halen", "Panama", "4"}
	band3 := []string{"The Lumineers", "Charlie Boy", "4"}
	csvWriter := csv.NewWriter(file)
	strWrite := [][]string{heading, band1, band2, band3}
	csvWriter.WriteAll(strWrite)
	csvWriter.Flush()
}
