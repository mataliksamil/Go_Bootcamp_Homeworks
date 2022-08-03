package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(power1(2, 5))
	fmt.Println(power2(2, 5))

	x := 2.5
	y := 10.8
	//fmt.Println(power1(x, y))
	fmt.Println(power2(int(x), int(y)), " : it is like 2^10")
	//fmt.Println(power2(2.5, 10.8))
	fmt.Println(math.Floor(4.5), math.Floor(5.8))
	// int() function also returns greatest integer value smaller than given float
	fmt.Println(power2(int(math.Floor(4.5)), int(math.Floor(5.8))), " : 4^5 = 1024")
}

// This function returns result of x powered by y operation
// which means result = x^y
func power1(x int, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}

// the only difference is x, y int vs x int, y int descriptions
func power2(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}
