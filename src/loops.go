package main

import "fmt"

func main() {
	// For conditional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n")
	// for while
	var counter int = 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	//var counterForever int = 0
	/*for {
		fmt.Println(counterForever)
		counterForever++
	}*/

	//For range
	names := []string{"mario", "luigi", "yoshi", "peach"}
	for index, value := range names {
		fmt.Printf("the position at index %v is: %v \n", index, value)
	}
}
