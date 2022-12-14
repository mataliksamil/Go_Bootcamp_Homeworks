// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

func main() {
	fmt.Println("Please Tell Me The Numbers ")
	// takas input
	reader := os.Args[1:]
	// an array from the input numbers
	var nums []int
	total := 0.0
	for _, i := range reader {
		num, err := strconv.Atoi(i)
		if err != nil { // in orfder to handle with not number input case
			num = 0
		}
		nums = append(nums, num) // appends slice converted inputs
		total += float64(num)    // in order to avg calculation
	}
	avg := total / float64(len(nums))
	fmt.Printf("Your nubers : %v\n", nums)
	fmt.Printf("Average : %.2f", avg)
}
