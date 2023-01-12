package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	fmt.Println("========\nTypes and Interfaces\n=======")

	// create a square using a new function
	fmt.Println(Rectangle{5, 8})

	sqr := newRectangle(6, 7)
	fmt.Printf("New square = %v\n", sqr)

	// create a rectangle directly (print with field names - %+v)
	rec := Rectangle{width: 3, height: 4}
	fmt.Printf("New rectangle = %+v\n", rec)

	circ := Circle{radius: 5}
	fmt.Printf("New Circle = %+v\n", circ)

	// print with JSON libraries
	// to-do: not sure why it's coming out {}
	u, err := json.Marshal(sqr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u))

	// try it a differerent way
	someRec := Rectangle{7, 7}
	fmt.Println(someRec)

	fmt.Println("\n***Print using JSON Marshal")
	// print using json marshal
	j, err := json.Marshal(someRec)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))

	// print shapes using interface
	fmt.Println("\n*** Print area of shapes using interface")
	printShapeInfo(sqr)
	printShapeInfo(rec)
	printShapeInfo(circ)
}

func printShapeInfo(s shape) {
	fmt.Printf("\nShape is %v", s.shapeType())
	fmt.Printf("\nShape area is %f\n", s.area())
}

// **********
// define the shape interface d
type shape interface {
	area() float64
	shapeType() string
}

// **********
// define the shape structures
type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// **********
// define the interface functions
func (rec Rectangle) area() float64 {
	return rec.width * rec.height
}

func (rec Rectangle) shapeType() string {
	return "Rectangle"
}

func (cir Circle) area() float64 {
	// return 2*r.width + 2*r.height
	return math.Pi * (cir.radius * cir.radius)
}

func (cir Circle) shapeType() string {
	return "Circle"
}

// **********
// define a function for construction
func newRectangle(height float64, width float64) *Rectangle {
	rec := Rectangle{height, width}
	return &rec
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
