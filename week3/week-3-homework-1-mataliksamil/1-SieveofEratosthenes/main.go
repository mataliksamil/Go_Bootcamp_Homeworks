package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	reader := os.Args[1:]             // input as a string
	num, _ := strconv.Atoi(reader[0]) // string to int conversation
	var primeSlice []int              // prime numbers that will be printed

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

	fmt.Printf(" Prime number list : \n %v", primeSlice)
}

func isPrime(num int) bool {

	var remainders []int // remainders slice
	for i := 2; i <= num/2+1; i++ {
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
