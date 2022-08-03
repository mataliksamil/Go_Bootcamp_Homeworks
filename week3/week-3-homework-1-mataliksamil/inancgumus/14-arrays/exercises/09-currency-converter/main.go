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
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the contants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go invalid
//     Invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 9.24 EUR
//     10.50 USD is 8.19 GBP
//     10.50 USD is 1186.71 JPY
//
//   go run main.go 1
//     1.00 USD is 0.88 EUR
//     1.00 USD is 0.78 GBP
//     1.00 USD is 113.02 JPY
// ---------------------------------------------------------

func main() {
	const (
		usdEur = iota
		usdGBP
		usdJPY
		usdTRY
	)
	// ratios array
	ratios := [...]float64{
		usdEur: 0.88,
		usdGBP: 0.78,
		usdJPY: 113.02,
		usdTRY: 17.49,
	}
	// takes input as string array
	usdAmount := os.Args[1:]
	// conver string array input to float64
	usdAmounFloat, err := strconv.ParseFloat(usdAmount[0], 64)
	if err != nil {
		fmt.Printf("ERROR : %s \n", err)
	}
	// print exchange rates
	fmt.Printf("%.2f USD is %.2f \tEUR\n", usdAmounFloat, usdAmounFloat*ratios[usdEur])
	fmt.Printf("%.2f USD is %.2f \tGBP\n", usdAmounFloat, usdAmounFloat*ratios[usdGBP])
	fmt.Printf("%.2f USD is %.2f \tJPY\n", usdAmounFloat, usdAmounFloat*ratios[usdJPY])
	fmt.Printf("%.2f USD is %.2f \tTRY\n", usdAmounFloat, usdAmounFloat*ratios[usdTRY])

}
