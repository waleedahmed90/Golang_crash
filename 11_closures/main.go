package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// Go can support anonymous functions which can form closures
	// We can define a fucntion inlien without having to name it

	fmt.Println("Closures")

	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

}
