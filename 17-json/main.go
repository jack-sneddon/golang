package main

import (
	"encoding/json"
	"fmt"
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

}
