package main

import "fmt"

func main() {

	switch module := 5 % 2; module { //The first part is the variable we declare, the second variable is the parsing of the variable in the switch sentence
	case 0:
		fmt.Println("It's even")
	default:
		fmt.Println("It's odd")
	}

	//Without a condition
	value := 200
	switch {
	case value > 100:
		fmt.Println("Value is hihger than 100")
	case value < 0:
		fmt.Println("Value is lower than 0")
	default:
		fmt.Println("No condition")

	}
}
