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

func main() {
	const (
		Nov = 11 - iota // 11 - 0 = 11
		Oct             // 11 - 1 = 10
		Sep             // 11 - 2 = 9
	)
	var a1 [32]int32
	var k int32 = 51515
	for i := 0; i < len(a1)-1; i++ {
		if i == 5 {
			a1[i] = k
		} else {
			a1[i] = int32(i)
			fmt.Println(&a1[i])
		}
	}

	fmt.Println("----------------------------")
	fmt.Println(k)
	fmt.Println("----------------------------")
	dumFunc(&a1)
	fmt.Println(Sep, Oct, Nov)
}
func dumFunc(p *[32]int32) {
	for i := 0; i < len(p)-1; i++ {
		fmt.Println(&p[i])
	}
}
