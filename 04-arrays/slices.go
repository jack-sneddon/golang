package main

import "fmt"

func playSlices() {

	fmt.Println("\n======Slices - Unlike arrays, slices are typed only by the elements they contain (not the number of elements)")
	mySlice := make([]string, 3)
	fmt.Println("mySlice (empty):", mySlice)

	// assignment
	mySlice[0] = "slc1"
	mySlice[1] = "slc2"
	mySlice[2] = "slc3"
	fmt.Println("mySlice (with stuff):", mySlice)

	fmt.Println("length of the slice:", len(mySlice))

	//append values, check length
	mySlice = append(mySlice, "slc4")
	mySlice = append(mySlice, "slc5", "slc6", "slc7")
	fmt.Println("mySlice = :", mySlice)
	fmt.Println("length of the slice:", len(mySlice))

	// create a slice of unknown size
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

	fmt.Println(cities)

	// copy a slice with assignment (shallow copy)
	sliceCopy1 := mySlice
	fmt.Println("Copy slice using assignment (shallow copy)")
	fmt.Printf("\nsource slice: %[1]v, address: %[1]p\n", mySlice)
	fmt.Printf("\ncopy slice:   %[1]v, address: %[1]p\n", sliceCopy1)

	sliceCopy1 = append(sliceCopy1, "slc8")
	fmt.Println("\nAppend to copy")
	fmt.Printf("\nsource slice: %[1]v, address: %[1]p\n", mySlice)
	fmt.Printf("\ncopy slice:   %[1]v, address: %[1]p\n", sliceCopy1)

	// copy a slice with copy function
	sliceCopy2 := make([]string, len(mySlice))
	copy(sliceCopy2, mySlice)

	fmt.Println("Copy slice using copy function")
	fmt.Printf("\nsource slice: %[1]v, address: %[1]p\n", mySlice)
	fmt.Printf("\ncopy slice:   %[1]v, address: %[1]p\n", sliceCopy2)

	// copy a slice with append
	var sliceCopy3 []string
	sliceCopy3 = append(sliceCopy3, mySlice...)

	fmt.Println("Copy slice using append function")
	fmt.Printf("\nsource slice: %[1]v, address: %[1]p\n", mySlice)
	fmt.Printf("\ncopy slice:   %[1]v, address: %[1]p\n", sliceCopy3)

	var sliceCopy4 []string
	sliceCopy4 = append(sliceCopy4, mySlice[1], sliceCopy1[3])
	fmt.Println("Copy slice using append function")
	fmt.Printf("\nsource slice: %[1]v, address: %[1]p\n", mySlice)
	fmt.Printf("\ncopy slice:   %[1]v, address: %[1]p\n", sliceCopy4)

	// *** use of make function
	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays
	slice2 := make([]int, 5) // len(a)=5
	fmt.Printf("\nslice2 len=%d cap=%d %v\n",
		len(slice2), cap(slice2), slice2)

	// specify the capcity
	slice3 := make([]int, 0, 6) // len(b)=0, cap(b)=5
	fmt.Printf("\nslice3 len=%d cap=%d %v\n",
		len(slice3), cap(slice3), slice3)

	slice4 := slice3[:2]
	fmt.Printf("\nslice4 len=%d cap=%d %v\n",
		len(slice4), cap(slice4), slice4)

	slice5 := slice4[1:5]
	fmt.Printf("\nslice5 len=%d cap=%d %v\n",
		len(slice5), cap(slice5), slice5)

}
