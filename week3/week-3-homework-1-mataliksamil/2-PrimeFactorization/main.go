package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := os.Args[1:]                  // takes input as a string
	inputVal, _ := strconv.Atoi(reader[0]) // string to int conversation
	var Pslice []int                       // in order to keep Sieve of Eratosthenes Primes
	var factorSlice []int                  //	in order to keep Prime factors
	Pslice = primeList(inputVal)           //	calls the code in the Sieve of Eratosthenes example
	fmt.Println(Pslice)
	quotient := inputVal

	for i := range Pslice {
		if quotient == 1 { // break if all quotients are found
			break
		}
		for i <= inputVal { //	Divide untill remainder is not equal to zero
			if quotient%Pslice[i] == 0 { // Check if it is divisable without remainder
				quotient = quotient / Pslice[i]
				factorSlice = append(factorSlice, Pslice[i])
			} else {
				break
			}
		}

	}

	fmt.Println(factorSlice)
}

func primeList(num int) []int {

	var primeSlice []int // prime numbers that will be printed

	if num >= 2 { // handles with x <=2 cases
		primeSlice = append(primeSlice, 2)
	} else {
		fmt.Printf("Prime number description is not valid if x<2")
	}

	for i := 2; i <= num; i++ { // Checks every smaller number if it is prime or not
		if isPrime(i) {
			primeSlice = append(primeSlice, i)
		}
	}

	//fmt.Printf(" Prime number list : \n %v", primeSlice)
	return primeSlice
}

func isPrime(num int) bool {

	var remainders []int            // remainders slice
	for i := 2; i <= num/2+1; i++ { // uses half of the numbers in order to reduce the operations
		remainders = append(remainders, num%i)
	}

	primeNum := true
	// checks remainder slice in order to detecet if there any
	// 0 remainder found
	for _, j := range remainders {
		if j == 0 {
			primeNum = false
		}
	}
	return primeNum
}
