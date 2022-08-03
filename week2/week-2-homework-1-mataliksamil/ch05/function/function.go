package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(power1(2, 5))
	//fmt.Println(power2(2, 5))

	x := 2.5
	y := 10.8
	//fmt.Println(power1(x, y))
	fmt.Println(power2(int(x), int(y)))
	//fmt.Println(power2(2.5, 10.8))
	fmt.Println(math.Floor(4.5), math.Floor(5.8))
	fmt.Println(power2(int(math.Floor(4.5)), int(math.Floor(5.8))))
}

func power1(x int, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}

func power2(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}
