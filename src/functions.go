package main

import "fmt"

func normalFunction(message string) {
	fmt.Print(message)
}

func tripleArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func tripleArgument2(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hello world\n")
	tripleArgument(3, 6, "Hello")
	tripleArgument2(4, 7, "Hi")
	var value int = returnValue(3)
	fmt.Println("The value is: ", value)

	var value1, value2 int = doubleReturn(5)
	fmt.Printf("The value 1 is: %d, the value 2 is: %d\n", value1, value2)
	//If we don't want to use one of the values, we use _ when we call the function

	var value3, _ int = doubleReturn(7)
	fmt.Printf("The value 3 is: %d\n", value3)
}
