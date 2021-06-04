package main

import (
	"fmt"
)

func main() {
	// Arrays
	// Arrays have to be of fixed length

	var fruitArr [2]string
	// Assign Values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(len(fruitArr))

	// Declaration plus assignment (Simulationeous)

	veggiesArray := [2]string{"carrot", "turnip"}

	fmt.Println(veggiesArray)

	//slices (when the length of the array/list is not known in advance)

	cars := []string{"toyota", "suzuki", "ford", "BMW", "Mercedes", "VolksWagon"}

	fmt.Println(cars)

	fmt.Println(cars[1:4]) //starts at 1 and stops at (4-1)

}
