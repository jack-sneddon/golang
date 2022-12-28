package main

import "fmt"

// Create a new type of 'teams'
// which is a slice of strings
type teams []string

func (t teams) print() {

	for i, team := range t {
		fmt.Println(i, team)
	}

}
