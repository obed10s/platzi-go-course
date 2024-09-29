package main

import "fmt"

type car struct {
	brand string
	year  int
	color string
	model string
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Another way to instance a struct
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
