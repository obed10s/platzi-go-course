package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Marco"] = 20
	fmt.Println(m)

	//Iterate map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//finding values
	value := m["Jose"]
	value_not_exist := m["Josep"]
	fmt.Println(value)
	fmt.Println(value_not_exist) //If the key does not exist in the dictionary, it will return 0
	value_not_exists, ok := m["Josep"]
	fmt.Println(value_not_exists, ok) // If the value does not exist in the dictionary, it will return false
}
