package main

import "fmt"

func main() {
	//We will declare a constant (const in C)
	const pi float64 = 3.141592654
	const pi2 = 3.1416
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	//We declare int variables
	base := 12
	var high int = 14
	var area int
	fmt.Println(base, high, area)

	//zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Calculate area of a square
	const base_square float64 = 10.0
	area_square := base_square * base_square
	fmt.Println("Area square: ", area_square)

}
