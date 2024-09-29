package mypackage

import "fmt"

//CarPublic with public access
type CarPublic struct {
	Brand string //lowercase variables are private, when the name of the variable starts with an uppercase, the variable will be public
	Year  int
	Model string
	Color string
}

type carPrivate struct {
	brand string
	year  int
	model string
	color string
}

//We are gonna print a message
func PrintMessage(text string) {
	fmt.Println(text)
}
