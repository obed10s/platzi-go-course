package main

import (
	"fmt"
	packagepointers "project-layout/src/package_pointers"
)

func main() {
	var myPC = packagepointers.Pc{Ram: 16, Disk: 1024, Brand: "Lenovo"}
	myPC.DuplicateDisk()
	myPC.DuplicateRAM()
	myPC.Ping()
	fmt.Println(myPC)
}
