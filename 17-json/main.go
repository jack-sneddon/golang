package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Active bool
	dob    string
}

func main() {
	fmt.Println("========\nJSON\n=======")

	//int
	someInt, _ := json.Marshal(1024)
	fmt.Println(string(someInt)) // 1024

	// boolean
	someBool, _ := json.Marshal(true)
	fmt.Println(string(someBool)) // true

	// string
	someString, _ := json.Marshal("Hi Mom")
	fmt.Println(string(someString)) // Hi mom

	// slice
	someSlice, _ := json.Marshal([]string{"biz", "bang", "bar"})
	fmt.Println(string(someSlice))

	// map
	someMap, _ := json.Marshal(map[string]int{"a": 1, "b": 2, "c": 3})
	fmt.Println(string(someMap))

	person, err := json.Marshal(Person{Name: "Bob Smith", Age: 25, Active: true, dob: "10-31-2023"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(person))

	// using struct tags
	fmt.Println("***Struct Tags to JSON")
	lotr := Book{}
	lotr.Title = "A Gentleman in Moscow"
	lotr.Author = "Tolken"
	lotr.Pages = 1000
	fmt.Println(lotr)

	gm := Book{}
	gm.Title = "Lord of the Rings"
	gm.Pages = 350
	fmt.Println(gm)

	out, err := json.MarshalIndent(gm, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\n", string(out))

}

// example of tags - metadata attached to fields
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}
