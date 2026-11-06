package structs

import (
	"fmt"
	"unsafe"
)

func showAllBasicTypes() {
	var age int32
	fmt.Printf("This is age of type int32 : %d bytes\n", unsafe.Sizeof(age))
	var population int64
	fmt.Printf("This is population of type int64 : %d bytes\n", unsafe.Sizeof(population))
	var qty int
	fmt.Printf("This is qty of type int : %d bytes\n", unsafe.Sizeof(qty))
}
