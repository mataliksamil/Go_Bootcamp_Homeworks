package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//array1()
	//array2()
	//array3()
	//array4()
	array5()
}

func array1() {
	var a [5]int
	fmt.Println(a)

	print(a)
	iterate(a)

	fmt.Println("Initializing")
	for i := 0; i < len(a); i++ {
		a[i] = i * i
	}

	print(a)
	iterate(a)

	var b [5]int = a // Copying an array
	fmt.Println("Are they equal? ", a == b)

	for i := 0; i < len(b); i++ {
		b[i] = i*i + 1
	}
	fmt.Println("Are they equal? ", a == b)
}

func array2() {
	var a [5]int = [5]int{1, 3, 5} // Partial assignment
	//var a1 [3]int = [...]int{1, 3, 5} // Partial assignment

	print(a)
	modify1(a)
	print(a)
	modify2(&a)
	print(a)
}

func array3() {
	var a1 [5]int = [5]int{0: 1, 1: 3, 4: 5} // Initializing using index
	print(a1)

	var a2 [10]int = [10]int{5: 1, 2} // Initializing using index
	fmt.Println(a2)

	var a3 [10]int = [10]int{5: 2} // Initializing using index
	fmt.Println(a3)
}

func array4() {
	//var a [-5]int
	//n := f()
	//var a [n]int

	//n := 5
	//var a1 [n]int // n must be constant
	//fmt.Println(a1)

	var a2 []int
	fmt.Println(a2, len(a2))

	//var a2 [5]int
	//fmt.Println(a2[-1])

	//i := f()
	//fmt.Println(a2[i])

}

func array5() {
	a1 := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println(a1)

	rand.Seed(time.Now().UnixNano()) // Initialize rand

	const columnCount = 5
	const rowCount = 4
	var a2 [rowCount][columnCount]int
	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			a2[i][j] = rand.Intn(10_000)
		}
	}
	fmt.Println(a2)

	a3 := a2[0]
	fmt.Println(a3)
	fmt.Println(a2[2][2])
}
func modify1(a [5]int) {
	fmt.Println("Modifying Elements")
	for i := 0; i < len(a); i++ {
		a[i] = 10 * i
	}
}

func modify2(a *[5]int) {
	fmt.Println("Printing Elements")
	for i := 0; i < len(a); i++ {
		a[i] = 10 * i
	}
}

func print(a [5]int) {
	fmt.Println("Printing Elements")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d: %d\n", i, a[i])
	}
}

func iterate(a [5]int) {
	fmt.Println("Iterating Elements")
	for i, e := range a {
		fmt.Printf("%d: %d\n", i, e)
	}
}

func f() int {
	return 5
}
