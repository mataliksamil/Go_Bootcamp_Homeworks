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
)

// ---------------------------------------------------------
// EXERCISE: Declare nil slices
//
//  1. Declare the following slices as nil slices:
//
//    1. The names of your friends (names slice)
//
//    2. The distances to locations (distances slice)
//
//    3. A data buffer (data slice)
//
//    4. Currency exchange ratios (ratios slice)
//
//    5. Up/Down status of web servers (alives slice)
//
//
//  2. Print their type, length and whether they're equal to nil value or not.
//
//
// EXPECTED OUTPUT
//  names    : []string 0 true
//  distances: []int 0 true
//  data     : []uint8 0 true
//  ratios   : []float64 0 true
//  alives   : []bool 0 true
// ---------------------------------------------------------

func main() {
	names := make([]string, 1)
	distances := make([]int, 1)
	data := make([]uint8, 1)
	ratios := make([]float64, 1)
	alives := make([]bool, 1)

	fmt.Printf("%T, %s, %t\n", names, names[0], names == nil)
	fmt.Printf("%T, %d, %t\n", distances, distances[0], distances == nil)
	fmt.Printf("%T, %d, %t\n", data, data[0], data == nil)
	fmt.Printf("%T, %f, %t\n", ratios, ratios[0], ratios == nil)
	fmt.Printf("%T, %t, %t\n", alives, alives[0], alives == nil)
}
