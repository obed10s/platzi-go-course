package packagepointers

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPc *Pc) DuplicateRAM() {
	myPc.Ram *= 2
}

func (myPc Pc) Ping() {
	fmt.Println(myPc.Brand, "Pong")
}

func (myPc *Pc) DuplicateDisk() {
	myPc.Disk *= 2
}

func (myPc Pc) String() string {
	return fmt.Sprintf("Brand: %s, Disk size: %d, RAM size: %d", myPc.Brand, myPc.Disk, myPc.Ram)
}
