package main

import "fmt"

func main() {
	//Sum
	x := 12
	y := 50
	result := x + y
	fmt.Println("The sum is: ", result)

	//Substraction
	result = y - x
	fmt.Println("The substraction is: ", result)

	//Product
	result = x * y
	fmt.Println("The product is: ", result)

	//Division
	result = y / x
	fmt.Println("The division is: ", result)

	//fmt package

	name := "Platzi"
	courses := 500

	fmt.Printf("%s has more than %d courses\n", name, courses)
	fmt.Printf("%v has more than %v courses\n", name, courses)

	message := fmt.Sprintf("%s has more than %d courses", name, courses)

	fmt.Println(message)

	//Data type
	fmt.Printf("name %T \n", name)
	fmt.Printf("courses %T \n", courses)

}
