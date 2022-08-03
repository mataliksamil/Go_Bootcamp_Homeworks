// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

func main() {
	fmt.Printf("\t First Name \t Last Name \t Nickname \n")
	fmt.Printf("\t =======================================\n")
	scientist := [5][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	fmt.Printf("\t %s \t %s \t %s \n", scientist[0][0], scientist[0][1], scientist[0][2])
	fmt.Printf("\t %s  \t %s \t %s \n", scientist[1][0], scientist[1][1], scientist[1][2])
	fmt.Printf("\t %s \t %s \t %s \n", scientist[2][0], scientist[2][1], scientist[2][2])
	fmt.Printf("\t %s  \t %s  \t %s \n", scientist[3][0], scientist[3][1], scientist[3][2])
	fmt.Printf("\t %s \t %s \t %s \n", scientist[4][0], scientist[4][1], scientist[4][2])
	fmt.Printf("\t ---------------------------------------\n")

}
