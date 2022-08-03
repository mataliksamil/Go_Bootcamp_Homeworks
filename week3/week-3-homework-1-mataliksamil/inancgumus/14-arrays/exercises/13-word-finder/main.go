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
// EXERCISE: Word Finder
//
//   Your goal is to search for the words inside the corpus.
//
//   Note: This exercise is similar to the previous word finder program:
//   https://github.com/inancgumus/learngo/tree/master/13-loops/10-word-finder-labeled-switch
//
//   1. Get the search query from the command-line (it can be multiple words)
//
//   2. Filter these words, do not search for them:
//      and, or, was, the, since, very
//
//      To do this, use an array for the filtered words.
//
//   3. Print the words found.
//
// RESTRICTION
//   + The search and the filtering should be case insensitive
//
// HINT
//   + strings.Fields function converts a given string to a slice.
//
//     You can find its example in the word finder program that I've mentioned
//     above.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me a word to search.
//
//   go run main.go and was
//
//   go run main.go AND WAS
//
//   go run main.go cat beginning
//     #2 : "cat"
//     #11: "beginning"
//
//   go run main.go Cat Beginning
//     #2 : "cat"
//     #11: "beginning"
// ---------------------------------------------------------

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {

	reader := os.Args[1:]
	splitCorpus := strings.Split(corpus, " ") // split the words to an array
	var isFound bool                          // in order to check if it is found in the current cell
	isAny := false                            // in order to keep track of if it is found in any cell

	/* 	filterList := [...]string{
	   		"and",
	   		"or",
	   		"was",
	   		"the",
	   		"since",
	   		"very",
	   	}

	   	// THIS IS SAME QUERY WTIH SEARCHING BUT SIMPLIFIED FOR FILTERING
	   	// I KNOW IT IS BRUTE FORCE PLEASE DONT HOLD YOUR RECCOMENDATIONS
	   	for _, seek := range filterList {
	   		isAny = false
	   		for i, word := range splitCorpus {
	   			// convert to lower case to check if it contains so that search become incase-sensitive
	   			isFound = strings.Contains(word, seek)
	   			if isFound == true {
	   				// THE GREAT FILTERING, BY ABSOLUTE BRUTE FORCE
	   				splitCorpus[i] = ""
	   			}
	   		}
	   	}
	*/

	for _, seek := range reader {
		isAny = false
		for i, word := range splitCorpus {
			// convert to lower case to check if it contains so that search become incase-sensitive
			isFound = strings.Contains(strings.ToLower(word), strings.ToLower(seek))
			if isFound == true {
				isAny = true
				// THE GREAT FILTERING, BY ABSOLUTE BRUTE FORCE
				if (word == "and") || (word == "was") || (word == "or") || (word == "the") || (word == "since") || (word == "very") {
					continue
				}
				fmt.Printf("\t #%d %q \n", i+1, word)
			}
		}
		if !isAny { // if the word is not in the slice
			fmt.Printf("\t The word : %q could not be found \n", seek)
		}
	}

}
