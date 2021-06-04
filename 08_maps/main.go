package main

import (
	"fmt"
)

func main() {
	fmt.Println("MAPS ARE KEY: VALUE PAIRS")

	emails := make(map[string]string) // name of the map[data_type_of_key]data_type_of_values

	// Assign KeyValues

	emails["butt"] = "butt@gmail.com"
	emails["haji"] = "haji@gmail.com"
	emails["hada"] = "hada@gmail.com"

	fmt.Println(len(emails))

	fmt.Println(emails)

	delete(emails, "haji")

	fmt.Println(emails)

	// Declaring map and adding values directly

	age := map[string]int{
		"john": 12,
		"paul": 15,
		"hada": 24,
	}

	fmt.Println(age)

}
