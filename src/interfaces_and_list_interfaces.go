package main

import "fmt"

type figures2D interface {
	area() float64
}

type square struct {
	base float64
}
type rectangle struct {
	base float64
	high float64
}

func (c square) area() float64 {
	return c.base * c.base
}

func (r rectangle) area() float64 {
	return r.base * r.high
}

func calculate(f figures2D) {
	fmt.Println("Area:", f.area())
}
func main() {
	var mySquare = square{base: 3}
	var myRectangle = rectangle{base: 6, high: 3}

	calculate(mySquare)
	calculate(myRectangle)

	//List interfaces
	myInterface := []interface{}{"Hello", 12, 5.1}
	fmt.Println(myInterface...)
}
