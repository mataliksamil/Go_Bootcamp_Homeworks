package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func remove(mySlice []int, index int) []int {
	copy(mySlice[index:], mySlice[index+1:]) // Shift a[i+1:] left one index.
	//fmt.Println(mySlice, " ----- -- - ")
	mySlice[len(mySlice)-1] = 0        // Erase last element (write zero value).
	mySlice = mySlice[:len(mySlice)-1] // Truncate slice.
	//fmt.Println(mySlice, " ----- -- - 2 ")
	return mySlice
}

func main() {
	fmt.Println("Helloworld")
	rand.Seed(time.Now().UnixNano())
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	strNum := ""
	index := 0

	// 1st digit
	index = rand.Intn(7) + 1            // sets the first index, +1 since it can not be 0
	strNum += strconv.Itoa(nums[index]) //	the secret number will be keeping as string
	nums = remove(nums, index)          // remove used number from the slice since digits should not be the same

	// 2nd digit
	index = rand.Intn(7)
	strNum += strconv.Itoa(nums[index])
	nums = remove(nums, index)

	// 3rd digit
	index = rand.Intn(6)
	strNum += strconv.Itoa(nums[index])
	nums = remove(nums, index)

	// 4th digit
	index = rand.Intn(5)
	strNum += strconv.Itoa(nums[index])
	nums = remove(nums, index)

	fmt.Println(strNum)
	var first string

	for a := 0; a < 10; a++ { // You have 10 turns for this game

		fmt.Printf("Please insert your 4 digit guess :")
		fmt.Scanln(&first) // first[0] - gives as ASCII decimal system 0,1,2.. - 48,49,50...
		guess := strings.Split(first, "")
		correct, misplaced := 0, 0
		for i, guessDigit := range guess {
			k := strings.IndexAny(strNum, guessDigit)
			if k != -1 { // if guess value's one digit appears in secretNum
				if strings.Index(strNum, guessDigit) == i { // if digits have the same index
					correct += 1
				} else {
					misplaced -= 1
				}
			}
		}

		// different print states for different cases
		if correct == 0 {
			if misplaced != 0 {
				fmt.Println(misplaced)
			} else {
				fmt.Println("")
			}
		} else {
			if misplaced != 0 {
				fmt.Println(correct, " ", misplaced)
			} else {
				fmt.Println(correct)
			}

		}
		if correct == 4 { // all correct
			fmt.Println(" CONGRATZZ !!!")
			break
		}
	}

}
