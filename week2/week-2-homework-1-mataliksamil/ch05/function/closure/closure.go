package main

import (
	"fmt"
)

func main() {

	printClosure := func(x int) {
		fmt.Printf("%d\n", x)
	}
	processEvenNumbers(10, printClosure)

	fmt.Println()

	sum := 0
	sumClosure := func(x int) {
		sum += x
	}
	processEvenNumbers(10, sumClosure)
	fmt.Println(sum)

	fmt.Println()

	product := 1
	productClosure := func(x int) {
		product *= x
	}
	processEvenNumbers(10, productClosure)
	fmt.Println(product)

	fmt.Println()

	var slice []int
	sliceClosure := func(x int) {
		slice = append(slice, x)
	}

	processEvenNumbers(10, sliceClosure)
	fmt.Println(slice)
}

func processEvenNumbers(n int, block func(int)) {
	for i := 2; i <= n; i += 2 {
		block(i)
	}
}
