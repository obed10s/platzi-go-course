package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pc{ram: 12, disk: 512, brand: "Asus"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)
}
