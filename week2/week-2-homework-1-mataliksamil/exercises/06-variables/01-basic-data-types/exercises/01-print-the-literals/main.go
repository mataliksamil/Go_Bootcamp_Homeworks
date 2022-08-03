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
// EXERCISE: Print the literals
//
//  1. Print a few integer literals
//
//  2. Print a few float literals
//
//  3. Print true and false bool constants
//
//  4. Print your name using a string literal
//
//  5. Print a non-english sentence using a string literal
//
// ---------------------------------------------------------

func main() {
	fmt.Println(132, 456, -5000, 999999999999999999)
	fmt.Println(3.25, -758.648, 123456789.65432198745612345678945612345678954655232854165146514654165146516516584168516516513513651)
	fmt.Println(true, false)
	fmt.Println("Muhammet Şamil Atalık")
	fmt.Println("مرحبا اسمي شامل") // arapçayı olduğu gibi (sağdan sola) bastırmadı
}
