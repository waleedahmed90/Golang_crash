package main

import "fmt"

func greeting(name string) string {
	return "World! " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Hello!", greeting("from golang"))
	fmt.Println("SUM = ", getSum(3, 4))
}
