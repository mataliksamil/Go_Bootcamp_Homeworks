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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Hipster's Love Bookstore Search Engine
//
//  Your goal is to allow people to search for books.
//
//  1. Create an array with the following book titles:
//      Kafka's Revenge
//      Stay Golden
//      Everythingship
//      Kafka's Revenge 2nd Edition
//
//  2. Get the search query from the command-line argument
//
//  3. Search for the books in the books array
//
//  4. When the program finds the book, print it.
//  5. Otherwise, print that the book doesn't exist.
//
//  6. Handle the errors.
//
// RESTRICTION:
//   + The search should be case insensitive.
//
// EXPECTED OUTPUT
//   go run main.go
//     Tell me a book title
//
//   go run main.go STAY
//     Search Results:
//     + Stay Golden
//
//   go run main.go sTaY
//     Search Results:
//     + Stay Golden
//
//   go run main.go "Kafka's Revenge"
//     Search Results:
//     + Kafka's Revenge
//     + Kafka's Revenge 2nd Edition
//
//   go run main.go void
//     Search Results:
//     We don't have the book: "void"
//
// HINTS:
//   + To find out whether a string contains another string value, you can use the strings.Contains function.
//   + To convert a string value to lowercase, you can use the strings.ToLower function.
//   + Check out the strings package for more information.
// ---------------------------------------------------------

func main() {
	// book list as an array
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}
	// takes input
	reader := os.Args[1:]

	var isFound bool // in order to check if it is found in the current cell
	isAny := false   // in order to keep track of if it is found in any cell
	fmt.Printf("\tSearch results:\n")

	//
	for i := range books {
		// convert to lower case to check if it contains so that search become incase-sensitive
		isFound = strings.Contains(strings.ToLower(books[i]), strings.ToLower(reader[0]))
		if isFound == true {
			fmt.Printf("\t + %s \n", books[i])
			isAny = true
		}
	}
	// if nothin is found
	if !isAny {
		fmt.Printf("\tWe don't have the book: %q", reader[0])
	}

}
