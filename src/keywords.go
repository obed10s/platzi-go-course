package main

import "fmt"

func main() {
	//Defer ->
	defer fmt.Println("Hello") //This line excecutes this line before finishing the function, so all the lines above weill run first

	fmt.Println("World")

	//Continue and break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		//continue
		if i == 2 {
			fmt.Println("It's 2")
			continue
		}

		//Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

}
