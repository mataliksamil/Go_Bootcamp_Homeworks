package main

import (
	"fmt"
)

func rotator(mySlice []int, step int, toRight bool) []int {
	var tempSlice []int
	if toRight == true { // if it is rotate to right
		for _, val := range mySlice[len(mySlice)-step:] {
			tempSlice = append(tempSlice, val) // appends last step number of items to the new slice
		}
		for _, val := range mySlice[:len(mySlice)-step] {
			tempSlice = append(tempSlice, val) // then appends rest of the slice on to new slice
		}
	}
	if toRight == false {
		for _, val := range mySlice[step:] {
			tempSlice = append(tempSlice, val)
		}
		for _, val := range mySlice[:step] {
			tempSlice = append(tempSlice, val)
		}
	}

	return tempSlice
}

func main() {
	fmt.Println("HelloWorld")
	mySlice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11} // example slice
	fmt.Println(mySlice1, "before the rotation ")
	// example function call
	mySlice2 := rotator(mySlice1, 5, false) //rotates mySlice1 to left for 5 steps
	fmt.Println(mySlice2)
}
