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
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const max = 5
	var uniques [max]int

	// It's harder to make a program dynamic using arrays
	// max, _ := strconv.Atoi(os.Args[1])
	// not allowed since arrays are static
	// var uniques [10]int //

	// is this BL BX kind of thing ??
loop:
	for found := 0; found < max; {
		n := rand.Intn(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques[found] = n
		found++
	}

	fmt.Println("\n\nuniques:", uniques)
}
