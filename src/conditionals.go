package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var value1, value2 int
	value1 = 4
	value2 = 9
	if value1 == value2 || value1 < value2 {
		fmt.Println("We enter in the conditional")
	}

	value, err := strconv.Atoi("43")
	if err != nil {
		log.Fatal(err) //If the function can't convert a string to a number, it blocks the program and finishes
	}
	fmt.Println(value)

}
