package main

import (
	"fmt"
)

func main() {
	fmt.Println("LOOPS")

	// Long Method

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++ //i = i + 1
	}

	// Short method

	for i := 0; i < 10; i++ {
		fmt.Printf("%dth iteration and value %d\n", i, i+1)
	}

	// FizzBuzz (every multiple of three print (Fizz) every multiple of 5 (Buzz) every multiple of 3 and 5 (FizzBuzz))

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 { //or if its divisible by 15 because GCD(3,5) = 1 => LCM(3,5)=15
			fmt.Printf("%d : FizzBuzz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d : Buzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d : Fizz\n", i)
		} else {
			fmt.Printf("%d : \n", i)
		}
	}

}
