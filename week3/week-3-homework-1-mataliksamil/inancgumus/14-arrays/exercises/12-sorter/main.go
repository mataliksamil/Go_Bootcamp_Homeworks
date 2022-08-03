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
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments is not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------

// swap function takes the cells will be swapped as pointers so does not need to return
func swap(i *int, j *int) {
	tempVal := *i // in order not to lose the data
	*i = *j
	*j = tempVal

}

func main() {
	reader := os.Args[1:]
	// an array from the input numbers
	var nums []int
	for _, i := range reader {
		num, err := strconv.Atoi(i)
		if err != nil { // in orfder to handle with not number input case
			num = 0
		}
		nums = append(nums, num) // appends slice converted inputs
	}

	for i := range nums {
		// -1 for since we are on a cel already -i for the cells before the current cell
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				swap(&nums[j], &nums[j+1])
			}
		}
	}
	fmt.Printf("%v", nums)

}
