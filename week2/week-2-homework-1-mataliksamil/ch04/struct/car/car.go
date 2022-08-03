package main

import "fmt"

// typical definition of a struct named car
type Car struct {
	make, model string
	year        int
	speed       int
	distance    int
	convertible bool
}

func main() {
	var c1 Car      // declaration
	fmt.Println(c1) // empty struct type of Car
	// typical assignments
	c1.make = "Mercedes"
	c1.model = "G"
	c1.year = 2022
	fmt.Println(c1)
}
