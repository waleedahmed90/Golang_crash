package main

import (
	"fmt"
)

func main() {
	fmt.Println("RANGE")

	ids := []int{3, 4, 5, 5, 632, 32, 234, 234, 24, 23, 5243, 523, 4}

	// Loop through ids using range

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println("SUM: ", sum)

	// Using Range with a map
	age := map[string]int{
		"john": 12,
		"paul": 15,
		"hada": 24,
	}
	// Key values
	for k, v := range age {
		fmt.Println(k, " : ", v)
	}

	// Keys only

	for k := range age {
		fmt.Println(k)
	}

	// values only

	for _, v := range age {
		fmt.Println(v)
	}
}
