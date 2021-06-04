package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)

	fmt.Printf("%T\n", a) //int
	fmt.Printf("%T\n", b) //*int

	//  Use * to read val from address
	fmt.Println(*b)  // 5 (value of a)
	fmt.Println(&a)  // &a means a's address
	fmt.Println(*&a) // *&a means value in a's address

	// Changing value with the pointer
	*b = 10

	fmt.Println(a) // 10

}
