package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())

	slice := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(sum(slice))
	fmt.Println(sum(slice...))

	fmt.Println(anotherSum(slice))
	//fmt.Println(anotherSum(1, 2, 3))

	fmt.Println(alternatingSum(false, 1, 2, 3, 4, 5, 6))
	fmt.Println(alternatingSum(true, 1, 2, 3, 4, 5, 6))

	//fmt.Printf("%T\n", sum)
	//fmt.Printf("%T\n", anotherSum)
}

func sum(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func anotherSum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

//func variadicFunction(values1 ...int, values2 ...int) {
//}

func alternatingSum(b bool, values ...int) int {
	sum := 0
	length := len(values)
	for i, value := range values {
		if !b && i == length-1 {
			continue
		} else {
			sum += value
		}
	}
	return sum
}

//func alternatingSum1(i int, values ...int) int {
//	sum := 0
//	length := len(values)
//	for i, value := range values {
//		if !b && i == length-1 {
//			continue
//		} else {
//			sum += value
//		}
//	}
//	return sum
//}
