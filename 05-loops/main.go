package main

import "fmt"

func main() {
	fmt.Println("\n======\nloops\n=====")

	// basic for loop
	fmt.Println("for:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
		if i == 5 {
			continue
		}
	}

	// while loop - do not have a dedicated keyword.  Use for
	fmt.Println("\nwhile:")
	keepLooping := true
	i := 0
	for keepLooping {
		if i > 5 {
			keepLooping = false
		}
		fmt.Printf("\n%d", i)
		i++
	}

	// range - iterate over objects of a structure
	fmt.Println("\nRange:")

	var numbers [3]int
	numbers[0] = 100
	numbers[1] = 101
	numbers[2] = 102

	for i, number := range numbers {
		fmt.Printf("\numbers[%d] = %d", i, number)
	}
	cars := []string{"honda", "toyota", "jeep", "audi"}
	for count, car := range cars {
		fmt.Printf("\ncar[%d] = %v", count, car)
	}

	usernamePass := map[string]string{"joe": "5f4dcc3b5aa765d61d8327deb882cf99", "jane": "7c6a180b36896a0a8c02787eeafb0e4c", "james": "6cb75f652a9b52798eb6cf2201057c73"}
	for user, pass := range usernamePass {
		fmt.Printf("\n%s : %s", user, pass)
	}

	for user := range usernamePass {
		fmt.Printf("\nuser: %v", user)
	}

}
