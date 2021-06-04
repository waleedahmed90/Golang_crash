package main

import (
	"fmt"
)

func main() {
	x := 15
	y := 6 + 9

	if x < y { // common practice is not to use parantheses around the condition
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x > y {
		fmt.Printf("%d is less than %d\n", y, x)
	} else {
		fmt.Printf("%d is equal to %d\n", x, y)
	}

	// color := "red"
	// color := "blue"
	color := "green"

	if color == "red" {
		fmt.Printf("color is %s\n", color)
	} else if color == "blue" {
		fmt.Printf("color is %s\n", color)
	} else {
		fmt.Printf("%s is not a recognized color\n", color)
	}

	// switch structure

	//name := "yakhbakh"
	//name := "haji"
	name := "inspector"

	switch name {
	case "yakhbakh":
		fmt.Printf("name: %s\n", name)
	case "haji":
		fmt.Printf("name: %s\n", name)
	default:
		fmt.Printf("name: %s is not recongnizable \n", name)
	}
	// && (AND)
	// || (OR)
	// !  (NOT)
}
