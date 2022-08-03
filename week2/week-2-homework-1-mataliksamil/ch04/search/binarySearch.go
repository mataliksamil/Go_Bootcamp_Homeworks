package main

import "fmt"

func main() {
	// ordered array of integers
	var integers = []int{-8, 0, 1, 3, 7, 18, 74, 101, 223, 395, 683, 911}

	var key = 3
	var found bool = BinarySearchIteratively(integers, key)
	fmt.Printf("Is %d in the array: %t \n", key, found)

	key = 11
	found = BinarySearchIteratively(integers, key)
	fmt.Printf("Is %d in the array: %t \n", key, found)

	key = 395
	found = BinarySearchRecursively(integers, key)
	fmt.Printf("Is %d in the array: %t \n", key, found)

	key = 200
	found = BinarySearchRecursively(integers, key)
	fmt.Printf("Is %d in the array: %t \n", key, found)
}

func BinarySearchIteratively(array []int, key int) bool {
	var found bool = false
	var length int = len(array)
	var start int = 0
	var end int = length - 1

	if key < array[start] {
		return false
	}
	if key > array[end] {
		return false
	}

	for found != true { // iterate until find
		var middle int = (end + start) / 2
		var element int = array[middle] // always checks middle of the remainder slice
		if key == element {
			found = true
			break // break case where key is found
		} else if key < element {
			end = middle - 1
		} else if key > element {
			start = middle + 1
		}

		if start > end {
			break // cannot find the key in the array
		}
	}
	return found
}

func BinarySearchRecursively(array []int, key int) bool {
	var found bool = false
	var length int = len(array)
	var start int = 0
	var end int = length - 1

	if key < array[start] { // key is not in the array
		return false
	}
	if key > array[end] { // key is not in the array 2
		return false
	}

	var middle int = (end + start) / 2 // always check nmiddle of the slice
	var element int = array[middle]
	if key == element { // key is found so return true
		return true
	} else {
		if key < element { // key is in the lower part of the array
			end = middle - 1
		} else if key > element { // key is in the upper half of the array
			start = middle + 1
		}
		slice := array[start : end+1]               // array cut into half
		found = BinarySearchRecursively(slice, key) // recursion if not found until it returns true
	}
	return found
}
