package main

import (
	"fmt"
)

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int 64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - an alias for uint8
	// rune - an alias for int32
	// float32 float64
	// complex64 complex128

	// Using var keyword
	var name string = "Insp"
	var age int32 = 29
	const isCool = true
	//isCool = false 	//is an error because reassignment/value change on a const type is not allowed

	fmt.Println(name, age)
	fmt.Printf("%T %T\n", name, age)
	fmt.Printf("%T\n", isCool)

}
