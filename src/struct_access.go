package main

import (
	"fmt"
	pk "project-layout/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Renault"
	myCar.Year = 2021
	myCar.Model = "Sandero"
	myCar.Color = "Black"

	fmt.Println(myCar)

	//var myOtherCar pk.carPrivate //We cannot use a private, we can use that struct in the same package
	//fmt.Println(myOtherCar)

	pk.PrintMessage("Hello world")
}
